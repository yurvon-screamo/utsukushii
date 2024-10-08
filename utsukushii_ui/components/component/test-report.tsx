"use client";

import * as React from "react";
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card";
import { ReportContent } from "@/lib/model";
import TestReportRow from "./test-report-row";
import useSWR from "swr";
import { DevModeCheck, DevModeCmdUri } from "@/lib/utils";
import { DevModeRun } from "./dev-mode-run";

export interface TestReportProps {
  content: ReportContent;
  reload: React.EventHandler<any>;
}

export default function TestReport(props: TestReportProps) {
  const isDevModeResponse = useSWR(DevModeCheck, fetch);
  if (isDevModeResponse.isLoading) return <div>loading...</div>;

  const content = props.content;
  const isDevMode =
    isDevModeResponse.data?.status && isDevModeResponse.data?.status < 300;

  return (
    <div className="z-10 w-full items-center justify-between font-mono text-sm lg:flex">
      <Card className="w-full">
        <CardHeader className="flex items-left justify-between">
          <div className="items-center grid grid-cols-3">
            <div className="flex items-center gap-2">
              <CalendarIcon
                onClick={() => fetch(DevModeCmdUri)}
                className="h-3 w-3 text-muted-foreground"
              />
              <span className="text-xs text-muted-foreground">
                {content.timestamp}
              </span>
            </div>
            <CardTitle>{content.title}</CardTitle>
            {isDevMode === true && <DevModeRun reload={props.reload} />}
          </div>
        </CardHeader>

        <CardContent>
          <div className="grid grid-cols-3 gap-4">
            <div className="flex items-center gap-2">
              <TestTubeIcon className="h-6 w-6 text-primary" />
              <div>
                <div className="text-lg font-medium">{content.total}</div>
                <div className="text-sm text-muted-foreground">Total Tests</div>
              </div>
            </div>
            <div className="flex items-center gap-2">
              <PercentIcon className="h-6 w-6 text-primary" />
              <div>
                <div className="text-lg font-medium">{content.coverage}%</div>
                <div className="text-sm text-muted-foreground">
                  Test Coverage
                </div>
              </div>
            </div>
            <div className="flex flex-col items-start gap-2">
              <div className="flex items-center gap-2">
                <CheckIcon className="h-5 w-5 text-green-500" />
                <div className="text-sm text-muted-foreground">
                  {content.success} Success
                </div>
              </div>
              <div className="flex items-center gap-2">
                <XIcon className="h-5 w-5 text-red-500" />
                <div className="text-sm text-muted-foreground">
                  {content.dropped} Dropped
                </div>
              </div>
              <div className="flex items-center gap-2">
                <PauseIcon className="h-5 w-5 text-yellow-500" />
                <div className="text-sm text-muted-foreground">
                  {content.skipped} Skipped
                </div>
              </div>
            </div>
          </div>
          <div className="mt-6">
            <TestReportRow records={content.tests}></TestReportRow>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}

function CalendarIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M8 2v4" />
      <path d="M16 2v4" />
      <rect width="18" height="18" x="3" y="4" rx="2" />
      <path d="M3 10h18" />
    </svg>
  );
}

function CheckIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M20 6 9 17l-5-5" />
    </svg>
  );
}

function PauseIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <rect x="14" y="4" width="4" height="16" rx="1" />
      <rect x="6" y="4" width="4" height="16" rx="1" />
    </svg>
  );
}

function PercentIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <line x1="19" x2="5" y1="5" y2="19" />
      <circle cx="6.5" cy="6.5" r="2.5" />
      <circle cx="17.5" cy="17.5" r="2.5" />
    </svg>
  );
}

function TestTubeIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M14.5 2v17.5c0 1.4-1.1 2.5-2.5 2.5h0c-1.4 0-2.5-1.1-2.5-2.5V2" />
      <path d="M8.5 2h7" />
      <path d="M14.5 16h-5" />
    </svg>
  );
}

function XIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M18 6 6 18" />
      <path d="m6 6 12 12" />
    </svg>
  );
}
