import React from "react";

const ProgramType: React.FC<{ Name: string; Type: string; Tag: string }> = ({
  Name,
  Tag,
  Type,
}) => {
  return (
    <div className="bg-gradient-to-br from-blue-50 to-indigo-50 p-4 rounded-lg shadow hover:shadow-md transition-all duration-200 border border-blue-100">
      <h2 className="font-semibold text-lg text-blue-700 mb-2">{Name}</h2>
      <div className="space-y-2">
        <div className="flex items-center">
          <span className="text-gray-600 w-16">Type:</span>
          <span className="text-gray-800 font-medium">{Type}</span>
        </div>
        <div className="flex items-center">
          <span className="text-gray-600 w-16">Tag:</span>
          <span className="bg-blue-100 text-blue-800 px-2 py-1 rounded-full text-sm font-medium">
            {Tag}
          </span>
        </div>
      </div>
    </div>
  );
};

export default ProgramType;
