import React from "react";
import { bpfprog } from "wailsjs/go/models";
import ProgramType from "./ProgramType";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";

const ProgList: React.FC<{ progList: bpfprog.ProgramInfo[] }> = ({
  progList,
}) => {
  return (
    <div>
      {progList?.length === 0 ? (
        <p className="text-gray-500 italic">No programs available</p>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
          {progList?.map((prog, index) => (
            <Sheet>
              <SheetTrigger>
                <ProgramType
                  key={index}
                  Name={prog.ProgramInfo?.Name || ""}
                  Tag={prog.ProgramInfo?.Tag || ""}
                  Type={prog.Type}
                />
              </SheetTrigger>
              <SheetContent className="min-w-screen text-black !animate-none !transition-none">
                <div className="min-w-screen">hi</div>
              </SheetContent>
            </Sheet>
          ))}
        </div>
      )}
    </div>
  );
};

export default ProgList;
