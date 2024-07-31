import { ReportContent } from "@/lib/model";
import { TestReport } from "@/components/component/test-report"
import fsPromises from 'fs/promises';
import path from 'path'

export default async function Home(props: any) {
  const filePath = path.join(process.cwd(), 'data.json');
  const jsonData = await fsPromises.readFile(filePath);
  const objectData : ReportContent = JSON.parse(jsonData.toString());

  return (
      <div className="z-10 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
        <TestReport content={objectData} />
      </div>
  );
}
