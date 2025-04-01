import { useEffect, useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { GetEbpfProgList } from "../wailsjs/go/main/App";
import { bpfprog, ebpf } from "../wailsjs/go/models";
import ProgramType from "./components/ProgramType";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import ProgList from "./components/ProgList";

function App() {
  const [progList, setProgList] = useState<bpfprog.ProgramInfo[]>();

  function fetchProgList() {
    GetEbpfProgList().then((progList) => {
      setProgList(progList);
    });
  }

  useEffect(() => {
    console.log("running");
    fetchProgList();
  }, []);

  return (
    <div id="App" className="p-6 bg-white rounded-lg shadow-lg">
      <ProgList progList={progList || []} />
    </div>
  );
}

export default App;
