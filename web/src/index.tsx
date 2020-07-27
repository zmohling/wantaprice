import React from 'react';
import ReactDOM from 'react-dom';
import * as serviceWorker from './serviceWorker';
import { Provider } from 'react-redux';
import store from './store/store';
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core';
import { blueGrey } from '@material-ui/core/colors/';
import Header from './components/header';
import Dashboard from './components/pages/dashboard';
import Register from './components/pages/register';

//set font for app
require('typeface-roboto');

const customTheme = createMuiTheme({
  palette: {
    primary: {
      main: blueGrey[700],
      light: '#718792',
      dark: '#1c313a'
    },
    secondary: {
      main: blueGrey[100],
      light: '#ffffff',
      dark: '#9ea7aa'
    }
  }
});

ReactDOM.render(
  <Provider store = {store}>
    <MuiThemeProvider theme={customTheme}>
      <Register/>
    </MuiThemeProvider>
  </Provider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
