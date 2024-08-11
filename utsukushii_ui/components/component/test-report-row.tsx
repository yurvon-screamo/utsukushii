"use client";

import * as React from "react";
import { useState } from "react";
import { Table, TableBody, TableRow, TableCell } from "@/components/ui/table";
import { State, TestRecord } from "@/lib/model";
import TestLog from "@/components/component/test-log";

export interface TestReportRowProps {
  records: TestRecord[];
}

export default function TestReportRow(props: TestReportRowProps) {
  const records = props.records;

  const [selectedTest, setSelectedTest] = useState<TestRecord | null>(null);
  const [expandedRow, setExpandedRow] = useState<number | null>(null);

  const handleRowClick = (rowIndex: number, test: TestRecord | null) => {
    setExpandedRow(rowIndex === expandedRow ? null : rowIndex);
    if (test && test.log) {
      setSelectedTest(selectedTest?.name === test.name ? null : test);
    }
  };

  return (
    <div className="justify-between font-mono lg:flex">
      <Table>
        <TableBody>
          {records.map((test, index) => (
            <React.Fragment key={index}>
              <TableRow
                onClick={() => handleRowClick(index, test)}
                className={`cursor-pointer ${
                  expandedRow === index ? "bg-muted" : ""
                }`}
              >
                <TableCell className="font-medium">{test.name}</TableCell>
                <TableCell
                  className={`font-medium ${
                    test.state === State.success
                      ? "text-green-500"
                      : test.state === State.dropped
                      ? "text-red-500"
                      : "text-yellow-500"
                  }`}
                >
                  {test.state}
                </TableCell>
                <TableCell>{test.duration}</TableCell>
              </TableRow>
              {expandedRow === index && test.tests && test.tests.length > 0 && (
                <TableRow>
                  <TableCell colSpan={3}>
                    <TestReportRow records={test.tests}></TestReportRow>
                  </TableCell>
                </TableRow>
              )}
            </React.Fragment>
          ))}
        </TableBody>
      </Table>
      {selectedTest?.log && <TestLog testName={selectedTest.name} testLog={selectedTest.log} />}
    </div>
  );
}
