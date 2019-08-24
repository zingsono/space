import React from 'react';
import ReactDOM from 'react-dom';

function Button(props){

    return <h1>Hello, world!</h1>
}


ReactDOM.render(
    <Button/>,
    document.getElementById('root')
);

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
