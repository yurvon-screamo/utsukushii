import TestReport from "@/components/component/test-report";
import { ContentUri } from "@/lib/utils";
import useSWR from "swr";

export default function Page() {
  const { data, error, isLoading, mutate } = useSWR(ContentUri, (u) =>
    fetch(u).then((c) => c.text())
  );

  if (error || !data) return <div>failed to load</div>;
  if (isLoading) return <div>loading...</div>;

  return (
      <TestReport reload={() => mutate()} content={JSON.parse(data)} />
  );
}
