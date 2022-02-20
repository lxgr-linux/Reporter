# Reporter
Tool to sum up (working) hours in markdown table.

## Installation
```shell
go install
```

## Usage
Enter the directory that contains your [`report.csv`](./report.csv) and run:
```shell
Reporter
```

It will output something like this:

| Start | End   | Time | Total |
| ----- | ----- | ---- | ----- |
| 10:30 | 12:45 | 2.25 |       |
| 23:00 |  0:30 | 1.50 |       |
|       |       |      | 3.75 |
