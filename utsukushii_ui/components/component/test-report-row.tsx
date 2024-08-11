"use client";

import * as React from "react";
import { useState } from "react";
import { Table, TableBody, TableRow, TableCell } from "@/components/ui/table";
import { State, TestRecord } from "@/lib/model";

export interface TestReportRowProps {
  records: TestRecord[];
}

export default function TestReportRow(props: TestReportRowProps) {
  const records = props.records;

  const [expandedRow, setExpandedRow] = useState<number | null>();
  const handleRowClick = (rowIndex: number) => {
    setExpandedRow(rowIndex === expandedRow ? null : rowIndex);
  };

  return (
    <Table>
      <TableBody>
        {records.map((test, index) => (
          <React.Fragment key={index}>
            <TableRow
              onClick={() => handleRowClick(index)}
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
  );
}
