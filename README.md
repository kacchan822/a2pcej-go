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

### Options

| Short | Long | Description |
| --- | --- | --- |
| `-m` | `--mode` | Conversion mode: `en` or `ja` (required) |
| `-d` | `--delimiter` | Use a custom delimiter |
| `-nd` | `--nodelimiter` | Use no delimiter |
| `-s` | `--sign` | Use a custom uppercase sign |
| `-ns` | `--nosign` | Use no uppercase sign |
| `-n` | `--num` | Convert digits to phonetic names |

## Testing

```sh
go test ./...
```
