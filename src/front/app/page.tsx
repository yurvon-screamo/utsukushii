import TestReport from "@/components/component/test-report";

export default async function Home(props: any) {
  
  return (  
    <div className="z-10 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
      <TestReport content={undefined!} />
    </div>
  );
}
