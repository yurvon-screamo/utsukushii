import TestReport, { TestReportProps } from "@/components/component/test-report";
import { GetServerSideProps, InferGetServerSidePropsType } from "next";
import { promises as fs } from 'fs';
import { ReportContent } from "@/lib/model";

export const getServerSideProps: GetServerSideProps<TestReportProps> = async () => {
  const file = await fs.readFile(process.cwd() + '/public/data.json', 'utf8');
  const content: ReportContent = JSON.parse(file);

  return {
    props: {
      content: content,
    },
  };
};

export default function Page(props: TestReportProps) {
  return (  
    <div className="z-10 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
      <TestReport content={props.content} />
    </div>
  );
}
