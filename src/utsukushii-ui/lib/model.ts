export interface ReportContent {
    title: string
    timestamp: string
    duration: string
    total: number
    coverage: number
    success: number
    dropped: number
    skipped: number
    groups: TestGroup[]
}

export interface TestGroup {
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
    success = "success",
    dropped = "dropped",
    skipped = "skipped",
} 
