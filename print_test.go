// print_test
package fmt

import (
	"strings"
	"sync"
	"testing"
)

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

	p := new(FmtProc)
	for i := 0; i < 2; i++ {
		p.Init(buf[:0])
		ret := Bprint(p, arg)
		if s := string(ret); strings.Compare(arg, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", arg, s)
		}
		if s := string(p.bufIn.bytes); strings.Compare(arg, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", arg, s)
		}
	}

	for i := 0; i < 2; i++ {
		p.Init(buf[:0])
		ret := Bprintf(p, "%s", arg)
		if s := string(ret); strings.Compare(arg, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", arg, s)
		}
		if s := string(p.bufIn.bytes); strings.Compare(arg, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", arg, string(p.bufIn.bytes))
		}
	}

	argLn := "hello word\n"
	for i := 0; i < 2; i++ {
		p.Init(buf[:0])
		ret := Bprintln(p, arg)
		if s := string(ret); strings.Compare(argLn, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", argLn, s)
		}
		if s := string(p.bufIn.bytes); strings.Compare(argLn, s) != 0 {
			t.Errorf("expect '%s', is not '%s' ", argLn, string(p.bufIn.bytes))
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
	p := new(FmtProc)
	for i := 0; i < b.N; i++ {
		p.Init(buf[:0])
		Bprintf(p, "%s", "hello word")
	}
}

var (
	printPool = &sync.Pool{New: func() interface{} { return new(FmtProc) }}
)

func BenchmarkBprintfWithPool(b *testing.B) {
	buf := make([]byte, 0, 128)
	for i := 0; i < b.N; i++ {
		p := printPool.Get().(*FmtProc)
		p.Init(buf[:0])
		Bprintf(p, "%s", "hello word")
		printPool.Put(p)
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintf("%s", "hello word")
	}
}
