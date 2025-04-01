import {useEffect, useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {GetEbpfProgList} from "../wailsjs/go/main/App";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    

    function greet() {
        GetEbpfProgList().then(progList => {
            setName(progList.join("--"))
        })
    }

    useEffect(()=> {
        greet();
    })

    return (
        <div id="App">
            prog list :
            {name}
        </div>
    )
}

export default App
