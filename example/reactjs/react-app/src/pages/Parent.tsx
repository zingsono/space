import React, {RefObject, useState} from 'react';

import {Child, Up} from './Child'

export interface Props {
    children: string;
    label: string;
}

export const Parent: React.FC<Props> = (props: Props) => {
    const [count, setCount] = useState(0);
    console.log('Parent.props = ', JSON.stringify(props))


    let plus = () => {
        setCount(count + 1)
    }

    let refChild: RefObject<Up> = React.createRef<Up>()

    return (
        <div className="Parent">
            <p>Count: {count}</p>
            <button onClick={plus}>Parent Button</button>
            <br/>
            <button onClick={()=>{refChild.current && refChild.current.updateUp(count)}}>Update Child Button</button>
            <div>Parent props.children = {props.children}</div>
            <Child {...props} onPlus={plus} ref={refChild}>childText</Child>
        </div>
    )
}
