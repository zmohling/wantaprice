import React from 'react';
import ReactDOM from 'react-dom';
import * as serviceWorker from './serviceWorker';
import { Provider } from 'react-redux';
import store from './store/store';
import Header from './components/header';
import Dashboard from './components/pages/dashboard';

//set font for app
require('typeface-roboto');

ReactDOM.render(
  <Provider store = {store}>
    <Header/>
    <Dashboard/>
  </Provider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
