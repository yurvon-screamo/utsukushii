import TestReport from "@/components/component/test-report";
import useSWR from "swr";

export default function Page() {
  const { data, error, isLoading } = useSWR("/utsukushii.json", (u) =>
    fetch(u).then((c) => c.text())
  );

  if (error || !data) return <div>failed to load</div>;
  if (isLoading) return <div>loading...</div>;

  return (
      <TestReport content={JSON.parse(data)} />
  );
}
