export interface ReportContent {
  timestamp: string
  title: string
  total: number
  coverage: number
  success: number
  dropped: number
  skipped: number
  groups: Group[]
}

export interface Group {
  name: string
  state: State
  duration: string
  tests: Test[]
}

export interface Test {
  name: string
  state: State
  duration: string
}

export enum State {
  success = 'success',
  dropped = 'dropped',
  skipped = 'skipped'
}
