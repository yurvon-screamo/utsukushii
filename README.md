# utsukushii

A set of utilities that will allow you to create a beautiful test report for your application.

> in development

## Install

With go:

```bash
go install github.com/yurvon-screamo/utsukushii@latest
```

## Usage

## Fast start for golang test

Install `go-junit-report`:

```bash
go install github.com/jstemmer/go-junit-report/v2@latest
```

Generate report:

```bash
go test -v 2>&1 ./... | go-junit-report -set-exit-code > junit.xml && utsukushii gen --junit ./junit.xml
```

Run serve:

```bash
utsukushii serve
```

Open <http://localhost:8080>

### Generate content file

Peak your test test report(s) (now support junit only), and generate content file:

```bash
utsukushii gen --junit ./my-junit1.xml --junit ./my-junit2.xml
```

if necessary add coverage, output path and report title params:

```bash
utsukushii gen --junit ./my-junit1.xml --junit ./my-junit2.xml -o my-utsukushii.json -t "my report" --coverage 65
```

### Serve

Run web-ui for given content file:

```bash
utsukushii serve
```

or add addr and content path params (on default ":8080" and "utsukushii.json"):

```bash
utsukushii serve --addr :18181 --content my-utsukushii.json
```

## Design

Main goal:

> Convert test result into beatufy report and present it

UJ:

1) get the test report(s) in
   * format: junit, ???
   * native run: dotnet, go, cypress, ???
2) generate utusukushi content file
3) run web server with utusukushi ui with content file  

![target](target.png)

```ascii
              1. merge and group multi reports                                                
              2. history rate from multi reports                                              
              3. different output format?                                                     
                                                                                              
                              ┌───────────────────────────────────┐                           
                              │                                   │                           
  junit──────────────────────►│                                   │                           
                              │                                   │                           
                              │                                   │                           
  vstest ────────────────────►│           generator               ├───────────────────► report
                              │                                   │                           
                              │                                   │                           
..allure─────────────────────►│                                   │                           
                              │                                   │                           
                              └───────────────────────────────────┘                           

other....                                                                                       
```

## TODO

* [X] Mvp on junit single input
* [X] Go serve
* [X] Mvp on junit multi input
* [ ] Docs write
* [ ] github action to fmt, release
* [ ] Native go runner
* [ ] Native c# runner
* [ ] Release v0.0.1
* [ ] github action release
