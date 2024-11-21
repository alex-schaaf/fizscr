# fizscr

A simple-as-possible CLI tool that scrapes the
[FIZ](https://my.sport.uni-goettingen.de/fiz/) gym website of the University of
Goettingen for its utilization between zero and one. The output is appended to a
file.

```bash
# build
go build -o fizscr ./cmd/scraper.go

# run
./fitzscr <outfile>

# outfile
cat <outfile>
2024-11-21T20:55:01Z	0.547826
```
