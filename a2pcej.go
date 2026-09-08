// Package a2pcej converts ASCII letters to English phonetic words or
// Japanese katakana letter names.
package a2pcej

import (
	"fmt"
	"strings"
)

// Options controls conversion. Empty Delimiter and Sign values are meaningful;
// use DefaultOptions to start with the language defaults.
type Options struct {
	Delimiter string
	Sign      string
	Num       bool
}

// Converter holds a reusable conversion configuration.
type Converter struct {
	lang      string
	delimiter string
	sign      string
	num       bool
	alphabet  []string
	numbers   []string
}

var englishAlphabet = []string{
	"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel",
	"India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa",
	"Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "Xray",
	"Yankee", "Zulu",
}

var japaneseAlphabet = []string{
	"エイ", "ビー", "シー", "ディー", "イー", "エフ", "ジー", "エイチ",
	"アイ", "ジェイ", "ケイ", "エル", "エム", "エヌ", "オー", "ピー",
	"キュー", "アール", "エス", "ティー", "ユー", "ヴィー", "ダブリュー",
	"エクス", "ワイ", "ゼット",
}

var englishNumbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var japaneseNumbers = []string{"ゼロ", "イチ", "ニイ", "サン", "ヨン", "ゴウ", "ロク", "シチ", "ハチ", "キュウ"}

// DefaultOptions returns the Python implementation's defaults for lang.
func DefaultOptions(lang string) (Options, error) {
	switch lang {
	case "en":
		return Options{Delimiter: "-", Sign: "(CAPS)"}, nil
	case "ja":
		return Options{Delimiter: "・", Sign: "（大文字）"}, nil
	default:
		return Options{}, fmt.Errorf("language %q is not supported", lang)
	}
}

// New creates a converter for "en" or "ja".
func New(lang string, options Options) (*Converter, error) {
	c := &Converter{lang: lang, delimiter: options.Delimiter, sign: options.Sign, num: options.Num}
	switch lang {
	case "en":
		c.alphabet, c.numbers = englishAlphabet, englishNumbers
	case "ja":
		c.alphabet, c.numbers = japaneseAlphabet, japaneseNumbers
	default:
		return nil, fmt.Errorf("language %q is not supported", lang)
	}
	return c, nil
}

// Convert converts each Unicode code point and joins the results with the delimiter.
// Only ASCII A-Z/a-z and, when enabled, ASCII digits are translated.
func (c *Converter) Convert(letters string) string {
	converted := make([]string, 0, len([]rune(letters)))
	for _, letter := range letters {
		switch {
		case letter >= 'A' && letter <= 'Z':
			converted = append(converted, c.alphabet[letter-'A']+c.sign)
		case letter >= 'a' && letter <= 'z':
			converted = append(converted, c.alphabet[letter-'a'])
		case c.num && letter >= '0' && letter <= '9':
			converted = append(converted, c.numbers[letter-'0'])
		default:
			converted = append(converted, string(letter))
		}
	}
	return strings.Join(converted, c.delimiter)
}

// ConvAL converts using the English defaults.
func ConvAL(letters string) string {
	opts, _ := DefaultOptions("en")
	c, _ := New("en", opts)
	return c.Convert(letters)
}

// ConvAK converts using the Japanese defaults.
func ConvAK(letters string) string {
	opts, _ := DefaultOptions("ja")
	c, _ := New("ja", opts)
	return c.Convert(letters)
}
