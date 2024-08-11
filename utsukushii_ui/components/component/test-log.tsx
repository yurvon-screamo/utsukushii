"use client";

import * as React from "react";

export interface TestLogProps {
  testName: string;
  testLog: string;
}

export default function TestLog(props: TestLogProps) {
  const fadeInStyle = {
    'white-space': "pre-wrap",
    transition: "opacity 500ms linear",
  };

  return (
    <div
      style={fadeInStyle}
      className="flex w-full text-xs"
    >
      {props.testLog}
    </div>
  );
}
