import React, { useState } from 'react';
import { Container, TextField, 
  CssBaseline, Button, Typography } from '@material-ui/core';
import { makeStyles, Theme } from '@material-ui/core/styles';

const useStyles = makeStyles((theme: Theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  title: {
    fontSize: 45,
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },

}));

const Register = (props: any) => {
  const styles = useStyles();
  const [name, setName] = useState('');
  const [phone, setPhone] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');

  const handleName = (e: any) => {
    setName(e.target.value);
  };

  const handlePhone = (e: any) => {
    setPhone(e.target.value);
  };

  const handleEmail = (e: any) => {
    setEmail(e.target.value);
  };
  
  const handlePassword = (e: any) => {
    setPassword(e.target.value);
  };

  const handleConfirmPassword = (e: any) => {
    setConfirmPassword(e.target.any);
  };

  async function createUser(){

    const url = 'https://api.wantaprice.com/v1/users';
    const body = {
      "displayName": name,
      "login": email,
      "phone": phone,
      "password": password
    };

    const res = await fetch(url, {
      method: 'POST',
      mode: 'no-cors',
      headers: {
        'Content-Type':'application/json',
      },
      body: JSON.stringify(body)
    });

    console.log(res);
    
  }

  let content = (
    <Container component="main" maxWidth="xs">
    <CssBaseline />
      <div className={styles.paper}>
        <Typography className={styles.title}>
          WantAPrice
        </Typography>
        <form onSubmit={(e:any) => {
          e.preventDefault();
          createUser();
        }}>
          <TextField
            margin="normal"
            variant="outlined"
            required
            fullWidth
            id="name"
            label="Name"
            name="name"
            onChange={handleName}
            autoFocus
          />
          <TextField
            margin="normal"
            variant="outlined"
            required
            fullWidth
            id="email"
            label="Email Address"
            name="email"
            autoComplete="email"
            onChange={handleEmail}
          />
          <TextField
            margin="normal"
            variant="outlined"
            required
            fullWidth
            id="phone"
            label="Phone Number "
            name="phone"
            autoComplete="phone"
            helperText="We use this to send you notifications!"
            onChange={handlePhone}
          />
          <TextField
            margin="normal"
            variant="outlined"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            onChange={handlePassword}
          />
          <TextField
            variant="outlined"
            required
            fullWidth
            name="confirm-password"
            label="Confirm Password"
            type="password"
            id="confirm-password"
            onChange={handleConfirmPassword}
          />

          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={styles.submit}
          >
            Register
          </Button>
        </form>

        <Typography>
          Already a user? <a href="#">Sign In</a>
        </Typography>
      </div>
    </Container>
  );
  
  return content;
}

export default Register;