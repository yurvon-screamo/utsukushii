import TestReport, { TestReportProps } from "@/components/component/test-report";
import raw from "@/data.json"

export default function Page() {
  const jsonView = JSON.stringify(raw)
  return (  
    <div className="z-10 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
      <TestReport content={JSON.parse(jsonView)} />
    </div>
  );
}
