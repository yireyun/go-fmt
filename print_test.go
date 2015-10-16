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
	argln := arg + "\n"
	//test: Fprint, Fprintf, Fprintln
	for i := 0; i < 2; i++ {
		buf = buf[:0]
		Fprint(&buf, arg)
		if s := string(buf); strings.Compare(arg, s) != 0 {
			t.Errorf("Fprint: expect '%s', is not '%s' ", arg, s)
		}
	}
	for i := 0; i < 2; i++ {
		buf = buf[:0]
		Fprintf(&buf, "%s", arg)
		if s := string(buf); strings.Compare(arg, s) != 0 {
			t.Errorf("Fprintf: expect '%s', is not '%s' ", arg, s)
		}
	}
	for i := 0; i < 2; i++ {
		buf = buf[:0]
		Fprintln(&buf, arg)
		if s := string(buf); strings.Compare(argln, s) != 0 {
			t.Errorf("Fprintln: expect '%s', is not '%s' ", argln, s)
		}
	}

	//test: Sprint, Sprintf, SprintLn
	for i := 0; i < 2; i++ {
		s := Sprint(arg)
		if strings.Compare(arg, s) != 0 {
			t.Errorf("Sprint: expect '%s', is not '%s' ", arg, s)
		}
	}
	for i := 0; i < 2; i++ {
		s := Sprintf("%s", arg)
		if strings.Compare(arg, s) != 0 {
			t.Errorf("Sprintf: expect '%s', is not '%s' ", arg, s)
		}
	}
	for i := 0; i < 2; i++ {
		s := Sprintln(arg)
		if strings.Compare(argln, s) != 0 {
			t.Errorf("SprintLn: expect '%s', is not '%s' ", argln, s)
		}
	}

	//test: Bprint, Bprintf, Bprintln
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
	for i := 0; i < 2; i++ {
		ret := Bprintln(buff[:0], arg)
		if s := string(ret); strings.Compare(argln, s) != 0 {
			t.Errorf("Bprintln: expect '%s', is not '%s' ", argln, s)
		}
	}

	p := new(Fmt)
	//Fmt.Fprint, Fmt.Fprintf, Fmt.Fprintln test
	for i := 0; i < 2; i++ {
		buf = buf[:0]
		p.Fprint(&buf, arg)
		if s := string(buf); strings.Compare(arg, s) != 0 {
			t.Errorf("Fprint: expect '%s', is not '%s' ", arg, s)
		}
	}
	for i := 0; i < 2; i++ {
		buf = buf[:0]
		p.Fprintf(&buf, "%s", arg)
		if s := string(buf); strings.Compare(arg, s) != 0 {
			t.Errorf("Fprintf: expect '%s', is not '%s' ", arg, s)
		}
	}
	for i := 0; i < 2; i++ {
		buf = buf[:0]
		p.Fprintln(&buf, arg)
		if s := string(buf); strings.Compare(argln, s) != 0 {
			t.Errorf("Fprintln: expect '%s', is not '%s' ", argln, s)
		}
	}

	//Fmt.Sprint, Fmt.Sprintf, Fmt.Sprintln test
	for i := 0; i < 2; i++ {
		s := p.Sprint(arg)
		if strings.Compare(arg, s) != 0 {
			t.Errorf("Sprint: expect '%s', is not '%s' ", arg, s)
		}
	}
	for i := 0; i < 2; i++ {
		s := p.Sprintf("%s", arg)
		if strings.Compare(arg, s) != 0 {
			t.Errorf("Sprintf: expect '%s', is not '%s' ", arg, s)
		}
	}
	for i := 0; i < 2; i++ {
		s := p.Sprintln(arg)
		if strings.Compare(argln, s) != 0 {
			t.Errorf("Sprintln: expect '%s', is not '%s' ", argln, s)
		}
	}

	//Fmt.Bprint, Fmt.Bprintf, Fmt.Bprintln test
	for i := 0; i < 2; i++ {
		ret := p.Bprint(buff[:0], arg)
		if s := string(ret); strings.Compare(arg, s) != 0 {
			t.Errorf("Bprint: expect '%s', is not '%s' ", arg, s)
		}
	}
	for i := 0; i < 2; i++ {
		ret := p.Bprintf(buff[:0], "%s", arg)
		if s := string(ret); strings.Compare(arg, s) != 0 {
			t.Errorf("Bprintf: expect '%s', is not '%s' ", arg, s)
		}
	}
	for i := 0; i < 2; i++ {
		ret := p.Bprintln(buff[:0], arg)
		if s := string(ret); strings.Compare(argln, s) != 0 {
			t.Errorf("Bprintln: expect '%s', is not '%s' ", argln, s)
		}
	}
}

var (
	fmtPool = sync.Pool{New: func() interface{} { return new(Fmt) }}
)

func BenchmarkFprint(b *testing.B) {
	var buf = buffer{}
	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		Fprint(&buf, "hello word")
	}
}
func BenchmarkFmtPoolFprint(b *testing.B) {
	var buf = buffer{}
	for i := 0; i < b.N; i++ {
		p := fmtPool.Get().(*Fmt)
		p.Fprint(&buf, "hello word")
		fmtPool.Put(p)
	}
}

func BenchmarkFmtFprint(b *testing.B) {
	var buf = buffer{}
	p := new(Fmt)
	for i := 0; i < b.N; i++ {
		p.Fprint(&buf, "hello word")
	}
}

func BenchmarkFprintln(b *testing.B) {
	var buf = buffer{}
	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		Fprintln(&buf, "hello word")
	}
}
func BenchmarkFmtPoolFprintLn(b *testing.B) {
	var buf = buffer{}
	for i := 0; i < b.N; i++ {
		p := fmtPool.Get().(*Fmt)
		p.Fprintln(&buf, "hello word")
		fmtPool.Put(p)
	}
}
func BenchmarkFmtFprintLn(b *testing.B) {
	var buf = buffer{}
	p := new(Fmt)
	for i := 0; i < b.N; i++ {
		p.Fprintln(&buf, "hello word")
	}
}

func BenchmarkFprintf(b *testing.B) {
	var buf = buffer{}
	for i := 0; i < b.N; i++ {
		buf = buf[:0]
		Fprintf(&buf, "%s", "hello word")
	}
}
func BenchmarkFmtPoolFprintf(b *testing.B) {
	var buf = buffer{}
	for i := 0; i < b.N; i++ {
		p := fmtPool.Get().(*Fmt)
		p.Fprintf(&buf, "%s", "hello word")
		fmtPool.Put(p)
	}
}
func BenchmarkFmtFprintf(b *testing.B) {
	var buf = buffer{}
	p := new(Fmt)
	for i := 0; i < b.N; i++ {
		p.Fprintf(&buf, "%s", "hello word")
	}
}

func BenchmarkBprint(b *testing.B) {
	buf := make([]byte, 0, 128)
	for i := 0; i < b.N; i++ {
		Bprint(buf[:0], "hello word")
	}
}
func BenchmarkFmtPoolBprint(b *testing.B) {
	buf := make([]byte, 0, 128)
	for i := 0; i < b.N; i++ {
		p := fmtPool.Get().(*Fmt)
		p.Bprint(buf[:0], "hello word")
		fmtPool.Put(p)
	}
}
func BenchmarkFmtBprint(b *testing.B) {
	buf := make([]byte, 0, 128)
	p := new(Fmt)
	for i := 0; i < b.N; i++ {
		p.Bprint(buf[:0], "hello word")
	}
}

func BenchmarkBprintln(b *testing.B) {
	buf := make([]byte, 0, 128)
	for i := 0; i < b.N; i++ {
		Bprintln(buf[:0], "hello word")
	}
}
func BenchmarkFmtPoolBprintLn(b *testing.B) {
	buf := make([]byte, 0, 128)
	for i := 0; i < b.N; i++ {
		p := fmtPool.Get().(*Fmt)
		p.Bprintln(buf[:0], "hello word")
		fmtPool.Put(p)
	}
}
func BenchmarkFmtBprintLn(b *testing.B) {
	buf := make([]byte, 0, 128)
	p := new(Fmt)
	for i := 0; i < b.N; i++ {
		p.Bprintln(buf[:0], "hello word")
	}
}

func BenchmarkBprintf(b *testing.B) {
	buf := make([]byte, 0, 128)
	for i := 0; i < b.N; i++ {
		Bprintf(buf[:0], "%s", "hello word")
	}
}
func BenchmarkFmtPoolBprintf(b *testing.B) {
	buf := make([]byte, 0, 128)
	for i := 0; i < b.N; i++ {
		p := fmtPool.Get().(*Fmt)
		p.Bprintf(buf[:0], "%s", "hello word")
		fmtPool.Put(p)
	}
}

func BenchmarkFmtBprintf(b *testing.B) {
	buf := make([]byte, 0, 128)
	p := new(Fmt)
	for i := 0; i < b.N; i++ {
		p.Bprintf(buf[:0], "%s", "hello word")
	}
}

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprint("hello word")
	}
}
func BenchmarkFmtPoolSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := fmtPool.Get().(*Fmt)
		p.Sprint("hello word")
		fmtPool.Put(p)
	}
}
func BenchmarkFmtSprint(b *testing.B) {
	p := new(Fmt)
	for i := 0; i < b.N; i++ {
		p.Sprint("hello word")
	}
}

func BenchmarkSprintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintln("hello word")
	}
}
func BenchmarkFmtPoolSprintLn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := fmtPool.Get().(*Fmt)
		p.Sprintln("hello word")
		fmtPool.Put(p)
	}
}
func BenchmarkFmtSprintLn(b *testing.B) {
	p := new(Fmt)
	for i := 0; i < b.N; i++ {
		p.Sprintln("hello word")
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintf("%s", "hello word")
	}
}
func BenchmarkFmtPoolSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := fmtPool.Get().(*Fmt)
		p.Sprintf("%s", "hello word")
		fmtPool.Put(p)
	}
}
func BenchmarkFmtSprintf(b *testing.B) {
	p := new(Fmt)
	for i := 0; i < b.N; i++ {
		p.Sprintf("%s", "hello word")
	}
}
