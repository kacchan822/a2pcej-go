package main

import (
	"flag"
	"fmt"
	"os"

	a2pcej "github.com/kacchan822/a2pcej-go"
)

func main() {
	mode := flag.String("mode", "", "conversion mode: en or ja (required)")
	flag.StringVar(mode, "m", "", "conversion mode: en or ja (required)")
	delimiter := flag.String("delimiter", "", "custom delimiter")
	flag.StringVar(delimiter, "d", "", "custom delimiter")
	nodelimiter := flag.Bool("nodelimiter", false, "use no delimiter")
	flag.BoolVar(nodelimiter, "nd", false, "use no delimiter")
	sign := flag.String("sign", "", "custom uppercase sign")
	flag.StringVar(sign, "s", "", "custom uppercase sign")
	nosign := flag.Bool("nosign", false, "use no uppercase sign")
	flag.BoolVar(nosign, "ns", false, "use no uppercase sign")
	num := flag.Bool("num", false, "convert digits")
	flag.BoolVar(num, "n", false, "convert digits")
	flag.Parse()

	if *mode == "" || flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: a2pcej -m en|ja [options] letters...")
		flag.PrintDefaults()
		os.Exit(2)
	}
	opts, err := a2pcej.DefaultOptions(*mode)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if *delimiter != "" {
		opts.Delimiter = *delimiter
	}
	if *nodelimiter {
		opts.Delimiter = ""
	}
	if *sign != "" {
		opts.Sign = *sign
	}
	if *nosign {
		opts.Sign = ""
	}
	opts.Num = *num
	c, err := a2pcej.New(*mode, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	for _, letters := range flag.Args() {
		fmt.Println(c.Convert(letters))
	}
}
