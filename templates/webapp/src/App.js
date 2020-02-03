import React from 'react';
import './App.css';
import {Technologies} from "./technologies/Technologies";

export function App() {
    return (
        <div className="app">
            <h2 className="title">Project Name</h2>
            <div>
                This project is generated with <b><a
                href="https://github.com/shpota/goxygen">goxygen</a></b>.

                <p/>The following list of technologies comes from
                a REST API call to the Go-based back end. Find
                the corresponding code
                in <code>webapp/src/technologies/Technologies.js
                </code> and <code>server/web/app.go</code>.
                <Technologies/>
            </div>
        </div>
    );
}
