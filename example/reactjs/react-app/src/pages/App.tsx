import React from 'react';
import logo from '../assets/logo.svg';
import './App.css';

import {Parent} from './Parent'

export default () => {
    let key = 'app.tsx key'
    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo"/>
                <p>
                    Edit <code>src/pages/App.tsx</code> and save to reload.
                </p>
                <Parent label={key}>11</Parent>
            </header>
        </div>
    );
}

