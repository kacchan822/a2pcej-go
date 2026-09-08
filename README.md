# a2pcej-go

Go implementation of [a2pcej](https://github.com/kacchan822/a2pcej), converting
ASCII alphabet letters to English phonetic code or Japanese katakana names.

```go
fmt.Println(a2pcej.ConvAL("Examples"))
// Echo(CAPS)-Xray-Alfa-Mike-Papa-Lima-Echo-Sierra

opts, _ := a2pcej.DefaultOptions("ja")
opts.Num = true
converter, _ := a2pcej.New("ja", opts)
fmt.Println(converter.Convert("A04"))
// エイ（大文字）・ゼロ・ヨン
```

## CLI

```sh
go run ./cmd/a2pcej -m en Examples004
go run ./cmd/a2pcej -m ja -n Examples004
go run ./cmd/a2pcej -m en -d ', ' -s '(CAPITAL)' Examples003
```

Options match the Python CLI: `-m/--mode`, `-d/--delimiter`,
`-nd/--nodelimiter`, `-s/--sign`, `-ns/--nosign`, and `-n/--num`.

Run tests with `go test ./...`.
