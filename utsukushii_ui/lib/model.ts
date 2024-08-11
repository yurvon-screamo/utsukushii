export interface ReportContent {
  title: string;
  timestamp: string;
  duration: string;
  total: number;
  coverage: number;
  success: number;
  dropped: number;
  skipped: number;
  tests: TestRecord[];
}

export interface TestRecord {
  name: string;
  state: State;
  duration: string;
  log: string | null;
  tests: TestRecord[];
}

export enum State {
  success = "success",
  dropped = "dropped",
  skipped = "skipped",
}
