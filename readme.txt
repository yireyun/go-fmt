c:/go/bin/go.exe test -test.bench=.* [C:/Go/src/fmt]
go1.11.10
goos: windows
goarch: amd64
pkg: fmt
BenchmarkSprintfPadding-8                  	10000000	       162 ns/op
BenchmarkSprintfEmpty-8                    	100000000	        16.7 ns/op
BenchmarkSprintfString-8                   	30000000	        50.5 ns/op
BenchmarkSprintfTruncateString-8           	20000000	        85.6 ns/op
BenchmarkSprintfSlowParsingPath-8          	30000000	        56.1 ns/op
BenchmarkSprintfQuoteString-8              	 5000000	       318 ns/op
BenchmarkSprintfInt-8                      	30000000	        35.6 ns/op
BenchmarkSprintfIntInt-8                   	20000000	        70.3 ns/op
BenchmarkSprintfPrefixedInt-8              	10000000	       132 ns/op
BenchmarkSprintfFloat-8                    	20000000	       117 ns/op
BenchmarkSprintfComplex-8                  	 5000000	       311 ns/op
BenchmarkSprintfBoolean-8                  	30000000	        44.7 ns/op
BenchmarkSprintfHexString-8                	10000000	       125 ns/op
BenchmarkSprintfHexBytes-8                 	10000000	       192 ns/op
BenchmarkSprintfBytes-8                    	 5000000	       312 ns/op
BenchmarkSprintfStringer-8                 	10000000	       226 ns/op
BenchmarkSprintfStructure-8                	 2000000	       690 ns/op
BenchmarkManyArgs-8                        	 5000000	       243 ns/op
BenchmarkFprintInt-8                       	10000000	       156 ns/op
BenchmarkFprintfBytes-8                    	 5000000	       238 ns/op
BenchmarkFprintIntNoAlloc-8                	10000000	       156 ns/op
BenchmarkFprintf-8                         	10000000	       152 ns/op
BenchmarkBprintfWithoutPool-8              	20000000	        86.1 ns/op
BenchmarkBprintfWithPool-8                 	10000000	       119 ns/op
BenchmarkSprintf-8                         	30000000	        63.5 ns/op
BenchmarkScanInts-8                        	    2000	    572177 ns/op
BenchmarkScanRecursiveInt-8                	      20	  75364675 ns/op
BenchmarkScanRecursiveIntReaderWrapper-8   	      20	  77460865 ns/op
PASS
ok  	fmt	48.931s
成功: 进程退出代码 0.
