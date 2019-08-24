import './style.css'

//import _ from 'lodash'
import printMe from './print.js';

import { cube } from './math.js';



function component() {
    let element = document.createElement('div');
    let btn = document.createElement('button');

    // Lodash（目前通过一个 script 脚本引入）对于执行这一行是必需的
    element.innerHTML = cube(5) // _.join(['Hello', 'webpack',cube(5)], ' ');

    btn.innerHTML = 'Click me and check the console!';
    btn.onclick = printMe;

    element.appendChild(btn);
    return element;
}

window.onload = ()=>{
    document.body.appendChild(component())
}


