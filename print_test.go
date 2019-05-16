// print_test
package fmt_test

import (
	. "fmt"
	"runtime"
	"strings"
	"sync"
	"testing"
	"unicode/utf8"
)

func init() {
	Printf("%v\n", runtime.Version())
}

type buffer []byte

func (b *buffer) Write(p []byte) (n int, err error) {
	*b = append(*b, p...)
	return len(p), nil
}

func (b *buffer) WriteString(s string) (n int, err error) {
	*b = append(*b, s...)
	return len(s), nil
}

func (b *buffer) WriteByte(c byte) error {
	*b = append(*b, c)
	return nil
}

func (bp *buffer) WriteRune(r rune) error {
	if r < utf8.RuneSelf {
		*bp = append(*bp, byte(r))
		return nil
	}

	b := *bp
	n := len(b)
	for n+utf8.UTFMax > cap(b) {
		b = append(b, 0)
	}
	w := utf8.EncodeRune(b[n:n+utf8.UTFMax], r)
	*bp = b[:n+w]
	return nil
}

func TestPrint(t *testing.T) {
	var buf = buffer{}
	arg := "hello word"

	for i := 0; i < 2; i++ {
		buf = buf[:0]
		Fprintf(&buf, "%s", "hello word")
		if strings.Compare(arg, string(buf)) != 0 {
			t.Errorf("expect '%s', is not '%s' ", arg, string(buf))
		}
	}

	for i := 0; i < 2; i++ {
		s := Sprintf("%s", arg)
		if strings.Compare(arg, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", arg, s)
		}
	}

	p := new(Fmt)
	for i := 0; i < 2; i++ {
		p.Init(buf[:0])
		ret := Bprint(p, arg)
		if s := string(ret); strings.Compare(arg, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", arg, s)
		}
		if s := string(p.Bytes()); strings.Compare(arg, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", arg, s)
		}
	}

	for i := 0; i < 2; i++ {
		p.Init(buf[:0])
		ret := Bprintf(p, "%s", arg)
		if s := string(ret); strings.Compare(arg, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", arg, s)
		}
		if s := string(p.Bytes()); strings.Compare(arg, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", arg, s)
		}
	}

	argLn := "hello word\n"
	for i := 0; i < 2; i++ {
		p.Init(buf[:0])
		ret := Bprintln(p, arg)
		if s := string(ret); strings.Compare(argLn, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", argLn, s)
		}
		if s := string(p.Bytes()); strings.Compare(argLn, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", argLn, s)
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

func BenchmarkBprintfWithoutPool(b *testing.B) {
	buf := make([]byte, 0, 128)
	p := new(Fmt)
	for i := 0; i < b.N; i++ {
		p.Init(buf[:0])
		Bprintf(p, "%s", "hello word")
	}
}

var (
	printPool = &sync.Pool{New: func() interface{} { return new(Fmt) }}
)

func BenchmarkBprintfWithPool(b *testing.B) {
	buf := make([]byte, 0, 128)
	for i := 0; i < b.N; i++ {
		p := printPool.Get().(*Fmt)
		p.Init(buf[:0])
		Bprintf(p, "%s", "hello word")
		printPool.Put(p)
	}
}

func BenchmarkSprintf(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Sprintf("%s", "hello word")
		}
	})
}
