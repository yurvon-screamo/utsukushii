"use client";

import * as React from "react";

export interface TestLogProps {
  testName: string;
  testLog: string;
}

export default function TestLog(props: TestLogProps) {
  return (
    <div className="flex w-full max-w-6xl text-sm text-muted-foreground">
      {props.testLog}
    </div>
  );
}
