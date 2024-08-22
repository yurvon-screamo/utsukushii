"use client";

import * as React from "react";

export interface TestLogProps {
  testName: string;
  testLog: string;
}

export default function TestLog(props: TestLogProps) {
  const { testLog } = props;

  return (
    <div
      className="bg-[#282c34] text-[#bbc2cf] px-4 py-2 text-xs font-mono w-full whitespace-pre transition-opacity ease-linear"
    >
      {testLog}
    </div>
  );
}
