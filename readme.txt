c:/go/bin/go.exe test -test.bench=.* [C:/Go/src/fmt]
go1.11.10
goos: windows
goarch: amd64
pkg: fmt
BenchmarkSprintfPadding-8                  	10000000	       191 ns/op
BenchmarkSprintfEmpty-8                    	100000000	        16.3 ns/op
BenchmarkSprintfString-8                   	20000000	        63.5 ns/op
BenchmarkSprintfTruncateString-8           	20000000	       107 ns/op
BenchmarkSprintfSlowParsingPath-8          	20000000	        74.4 ns/op
BenchmarkSprintfQuoteString-8              	 5000000	       323 ns/op
BenchmarkSprintfInt-8                      	20000000	        61.9 ns/op
BenchmarkSprintfIntInt-8                   	20000000	        94.6 ns/op
BenchmarkSprintfPrefixedInt-8              	10000000	       216 ns/op
BenchmarkSprintfFloat-8                    	10000000	       124 ns/op
BenchmarkSprintfComplex-8                  	 5000000	       365 ns/op
BenchmarkSprintfBoolean-8                  	20000000	        61.9 ns/op
BenchmarkSprintfHexString-8                	 5000000	       398 ns/op
BenchmarkSprintfHexBytes-8                 	 3000000	       552 ns/op
BenchmarkSprintfBytes-8                    	 3000000	       542 ns/op
BenchmarkSprintfStringer-8                 	 5000000	       238 ns/op
BenchmarkSprintfStructure-8                	 2000000	       923 ns/op
BenchmarkManyArgs-8                        	 5000000	       381 ns/op
BenchmarkFprintInt-8                       	10000000	       201 ns/op
BenchmarkFprintfBytes-8                    	 5000000	       264 ns/op
BenchmarkFprintIntNoAlloc-8                	10000000	       207 ns/op
BenchmarkFprintf-8                         	10000000	       178 ns/op
BenchmarkBprintfWithoutPool-8              	30000000	        59.7 ns/op
BenchmarkBprintfWithPool-8                 	20000000	       107 ns/op
BenchmarkSprintf-8                         	20000000	        91.2 ns/op
BenchmarkScanInts-8                        	    3000	    602219 ns/op
BenchmarkScanRecursiveInt-8                	      20	  70400140 ns/op
BenchmarkScanRecursiveIntReaderWrapper-8   	      20	  77217070 ns/op
PASS
ok  	fmt	54.294s
成功: 进程退出代码 0.
