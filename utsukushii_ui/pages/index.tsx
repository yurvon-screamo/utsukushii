import TestReport from "@/components/component/test-report";
import useSWR from "swr";

export default function Page() {
  const { data, error, isLoading } = useSWR("/utsukushii.json", (u) =>
    fetch(u).then((c) => c.text())
  );

  if (error || !data) return <div>failed to load</div>;
  if (isLoading) return <div>loading...</div>;

  return (
    <div className="z-10 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
      <TestReport content={JSON.parse(data)} />
    </div>
  );
}
