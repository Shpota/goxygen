import React from 'react';
import './App.css';
import logo from './logo.svg';
import {Tech} from "./tech/Tech";

export function App() {
    return (
        <div className="app">
            <h2 className="title">project-name</h2>
            <div className="logo"><img src={logo} height="150px" alt="logo"/></div>
            <div>
                This project is generated with <b><a
                href="https://github.com/shpota/goxygen">goxygen</a></b>.
                <p/>The following list of technologies comes from
                a REST API call to the Go-based back end. Find
                and change the corresponding code
                in <code>webapp/src/tech/Tech.js
                </code> and <code>server/web/app.go</code>.
                <Tech/>
            </div>
        </div>
    );
}
