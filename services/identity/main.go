package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/oklog/pkg/group"
)

func main() {
	// Create a single logger, which we'll use and give to other components.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var db *sql.DB
	{
		db = getDBConnection(logger)
		//db.SetLogger(gormLogWrapper{l: logger})
	}
	defer db.Close()

	rep, _ := sql.Open("postgres", os.Getenv("DB_CONN"))

	var (
		service = NewService(rep)
	)

	mux := http.NewServeMux()

	mux.Handle("/v1/users", MakeHandler(service, logger))

	var g group.Group
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", ":8080")
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", ":8080")
			return http.Serve(httpListener, mux)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		// // The gRPC listener mounts the Go kit gRPC server we created.
		// grpcListener, err := net.Listen("tcp", *grpcAddr)
		// if err != nil {
		// 	logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		// 	os.Exit(1)
		// }
		// g.Add(func() error {
		// 	logger.Log("transport", "gRPC", "addr", *grpcAddr)
		// 	// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
		// 	// the here demonstrated zipkin tracing middleware.
		// 	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
		// 	addpb.RegisterAddServer(baseServer, grpcServer)
		// 	return baseServer.Serve(grpcListener)
		// }, func(error) {
		// 	grpcListener.Close()
		// })
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}

func getDBConnection(logger log.Logger) *sql.DB {
	connStr, set := os.LookupEnv("DB_CONN")
	if set == false {
		logger.Log("exit", errors.New("DB connection string environment variable unset"))
		os.Exit(-1)
	}

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		level.Error(logger).Log("exit", err)
		panic("Failed to connect DB")
	}
	logger.Log("success", "DB connection established")

	return db
}
