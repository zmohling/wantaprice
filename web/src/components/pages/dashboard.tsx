import React from 'react';
import { 
  Container,
  Grid,
  TextField,
  makeStyles,
  Theme,
  Divider,
  withStyles,
  InputAdornment,
  IconButton, 
  Typography,
  Paper} from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';

const CustomTextField = withStyles({
  root: {
    '& .MuiOutlinedInput-root': {
      borderRadius: 50
    },
  },
})(TextField);

const useStyles = makeStyles((theme: Theme) => ({
  root: {
    flexGrow: 1,
    '&.MuiOutlinedInput-root':{
      borderRadius: 50
    }
  },
  search: {
    marginTop: '20px',
    marginBottom: '20px',
  },
  spacer: {
    height: '20px'
  },
  dashboardTitle: {
    fontSize: 20,
    margin: '10px'
  },
  card: {
    height: '200px',
  },
  toolbar: theme.mixins.toolbar,
}));

const Dashboard = (props: any) => {
  const styles = useStyles();

  let content = (
    <div className={styles.root}>
      <div className={styles.toolbar}/>
      <Container component="main" maxWidth="lg" >
        <Grid container justify="center" spacing={1}>
          <Grid item xs={10}>
            <CustomTextField
              className={styles.search}
              InputProps={{
                style: {
                  fontSize: 20,
                  color: 'rgb(70, 50, 70)'
                },
                startAdornment: (
                  <InputAdornment position="start">
                    <IconButton size="medium">
                      <SearchIcon fontSize="large"/>
                    </IconButton>
                  </InputAdornment>
                )
              }}
              placeholder="Search for a Product"
              id="search"
              name="search"
              fullWidth
              variant="outlined"
            />
          </Grid>
          
        </Grid>
        <Divider/>
        <Typography className={styles.dashboardTitle}>Your Dashboard</Typography>

        <Grid container spacing={2}>
          <Grid item xs ={3}>
            <Paper elevation={3} className={styles.card}/>
          </Grid>
          <Grid item xs ={3}>
            <Paper elevation={3} className={styles.card}/>
          </Grid>
          <Grid item xs ={3}>
            <Paper elevation={3} className={styles.card}/>
          </Grid>
          <Grid item xs ={3}>
            <Paper elevation={3} className={styles.card}/>
          </Grid>
          <Grid item xs ={3}>
            <Paper elevation={3} className={styles.card}/>
          </Grid>
          <Grid item xs ={3}>
            <Paper elevation={3} className={styles.card}/>
          </Grid>


        </Grid>
      </Container>
    </div>
  );
  return content;
}

export default Dashboard;