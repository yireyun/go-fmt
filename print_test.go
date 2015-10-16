// print_test
package fmt

import (
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	var buf = buffer{}
	arg := "hello word"

	for i := 0; i < 2; i++ {
		buf = buf[:0]
		Fprintf(&buf, "%s", "hello word")
		if strings.Compare(arg, string(buf)) != 0 {
			t.Errorf("Fprintf: expect '%s', is not '%s' ", arg, string(buf))
		}
	}

	for i := 0; i < 2; i++ {
		s := Sprintf("%s", arg)
		if strings.Compare(arg, s) != 0 {
			t.Errorf("Sprintf: expect '%s', is not '%s' ", arg, s)
		}
	}

	var buff = make([]byte, 0, 64)
	for i := 0; i < 2; i++ {
		ret := Bprint(buff[:0], arg)
		if s := string(ret); strings.Compare(arg, s) != 0 {
			t.Errorf("Bprint: expect '%s', is not '%s' ", arg, s)
		}
	}

	for i := 0; i < 2; i++ {
		ret := Bprintf(buff[:0], "%s", arg)
		if s := string(ret); strings.Compare(arg, s) != 0 {
			t.Errorf("Bprintf: expect '%s', is not '%s' ", arg, s)
		}
	}

	argLn := "hello word\n"
	for i := 0; i < 2; i++ {
		ret := Bprintln(buff[:0], arg)
		if s := string(ret); strings.Compare(argLn, s) != 0 {
			t.Errorf("Bprintln: expect '%s', is not '%s' ", argLn, s)
		}
	}
}

func BenchmarkFprintf(b *testing.B) {
	var buf = buffer{}
	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		Fprintf(&buf, "%s", "hello word")
	}
}

func BenchmarkBprintf(b *testing.B) {
	buf := make([]byte, 0, 128)
	for i := 0; i < b.N; i++ {
		Bprintf(buf[:0], "%s", "hello word")
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintf("%s", "hello word")
	}
}
