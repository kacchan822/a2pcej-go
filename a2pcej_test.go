package a2pcej

import "testing"

func TestConvenienceFunctions(t *testing.T) {
	if got, want := ConvAL("Examples002"), "Echo(CAPS)-Xray-Alfa-Mike-Papa-Lima-Echo-Sierra-0-0-2"; got != want {
		t.Fatalf("ConvAL = %q, want %q", got, want)
	}
	if got, want := ConvAK("HogE"), "エイチ（大文字）・オー・ジー・イー（大文字）"; got != want {
		t.Fatalf("ConvAK = %q, want %q", got, want)
	}
}

func TestOptions(t *testing.T) {
	opts := Options{Delimiter: ", ", Sign: "(CAPITAL)", Num: true}
	c, err := New("en", opts)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := c.Convert("A04#"), "Alfa(CAPITAL), zero, four, #"; got != want {
		t.Fatalf("Convert = %q, want %q", got, want)
	}
}

func TestUnicodeIsPreservedByRune(t *testing.T) {
	opts, _ := DefaultOptions("ja")
	c, _ := New("ja", opts)
	if got, want := c.Convert("aあ"), "エイ・あ"; got != want {
		t.Fatalf("Convert = %q, want %q", got, want)
	}
}

func TestUnsupportedLanguage(t *testing.T) {
	if _, err := DefaultOptions("fr"); err == nil {
		t.Fatal("expected error")
	}
	if _, err := New("fr", Options{}); err == nil {
		t.Fatal("expected error")
	}
}
