import React from 'react';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Redirect,
} from 'react-router-dom';
import { useSelector } from 'react-redux';

import Login from './pages/login';
import Register from './pages/register';
import Header from './header';
import Dashboard from './pages/dashboard';
import { useCookies } from 'react-cookie';

export default function Routing() {
    const loggedIn = useSelector(state => state.auth.loggedIn);
    const [cookies, setCookie] = useCookies(['name']);
    setCookie('name', 'test', { path: '/login' });

    return (
        <Router>
            <Switch>
                <Route exact path="/">
                    { (loggedIn || cookies.name) ? <div><Header /> <Dashboard /></div> : <Redirect to="/login" /> }
                </Route>
                <Route path="/login">
                    { (loggedIn || cookies.name) ? <Redirect to="/" /> : <Login /> }
                </Route>
                <Route path="/register">
                    <Register />
                </Route>
            </Switch>
        </Router>
    );
}