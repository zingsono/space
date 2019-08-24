import React from 'react';
import ReactDOM from 'react-dom';

function Button(props) {

    return React.createElement(
        'h1',
        null,
        'Hello, world!'
    );
}

ReactDOM.render(React.createElement(Button, null), document.getElementById('root'));

/*

function f() {
    let div = document.createElement("div")
    div.setAttribute("id","root")
    document.body.appendChild(div)
    return div
}

f()

ReactDOM.render(
    <h1>Hello, world!</h1>,
    document.getElementById('root')
);
*/