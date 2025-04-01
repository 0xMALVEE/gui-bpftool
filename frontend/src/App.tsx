import {useEffect, useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {GetEbpfProgList} from "../wailsjs/go/main/App";
import { ebpf } from '../wailsjs/go/models';

function App() {
    const [progList, setProgList] = useState<ebpf.ProgramInfo[]>()
    

    function greet() {
        GetEbpfProgList().then(progList => {
            setProgList(progList)
        })
    }

    useEffect(()=> {
        greet();
    })

    return (
        <div id="App" className='text-red-500'>
            prog list :
            {progList?.map(prog=> <div>
                Name: {prog.Name}
                Type: {prog.Type}
                Tag: {prog.Type}
            </div>)}
        </div>
    )
}




export default App
