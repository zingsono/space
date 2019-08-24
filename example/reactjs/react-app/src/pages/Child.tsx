import React, {FC, forwardRef, Ref, useImperativeHandle, useState} from 'react';

//import {MyContext} from './Parent'

export interface Props {
    children: string;
    onPlus: Function;
    onChildMethod?: Function;
    ref: Ref<Up>;
}

export interface Up {
    updateUp: Function;
}

export const Child: FC<Props> = forwardRef((props: Props, ref: Ref<Up>) => {
    console.log('ChildProps=', JSON.stringify(props))
    const [up, setUp] = useState(0);
    useImperativeHandle(ref, () => ({
        updateUp: (p: number) => {
            console.log('updateUp', p)
            setUp(p)
        }
    }))

    return (
        <div className="Child">
            <button onClick={() => {
                props.onPlus()
            }}>Child Button
            </button>
            <p>Child props.children ={props.children}</p>
            <p>up={up}</p>
        </div>
    )
})
