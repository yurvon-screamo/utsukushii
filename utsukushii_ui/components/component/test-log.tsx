"use client";

import * as React from "react";

export interface TestLogProps {
  testName: string;
  testLog: string;
}

export default function TestLog(props: TestLogProps) {
  const fadeInStyle = {
    "white-space": "pre-wrap",
    transition: "opacity 500ms linear",
  };

  return (
    <div style={fadeInStyle} className="border-l px-2 flex text-xs">
      {props.testLog}
    </div>
  );
}
