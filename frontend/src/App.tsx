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
        console.log("running")
        greet();
    }, [])

    return (
        <div id="App" className="p-6 bg-white rounded-lg shadow-lg">
        <h1 className="text-3xl font-bold mb-6 text-gray-800">Program List</h1>
        
        {progList?.length === 0 ? (
          <p className="text-gray-500 italic">No programs available</p>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {progList?.map((prog, index) => (
              <div 
                key={index} 
                className="bg-gradient-to-br from-blue-50 to-indigo-50 p-4 rounded-lg shadow hover:shadow-md transition-all duration-200 border border-blue-100"
              >
                <h2 className="font-semibold text-lg text-blue-700 mb-2">{prog.Name}</h2>
                <div className="space-y-2">
                  <div className="flex items-center">
                    <span className="text-gray-600 w-16">Type:</span>
                    <span className="text-gray-800 font-medium">{prog.Type}</span>
                  </div>
                  <div className="flex items-center">
                    <span className="text-gray-600 w-16">Tag:</span>
                    <span className="bg-blue-100 text-blue-800 px-2 py-1 rounded-full text-sm font-medium">{prog.Type}</span>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    )
}




export default App
