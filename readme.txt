c:/go/bin/go.exe test -test.bench=.* [C:/Go/src/fmt]
go1.10.8
goos: windows
goarch: amd64
pkg: fmt
BenchmarkSprintfPadding-8                  	10000000	       188 ns/op
BenchmarkSprintfEmpty-8                    	100000000	        15.9 ns/op
BenchmarkSprintfString-8                   	30000000	        64.7 ns/op
BenchmarkSprintfTruncateString-8           	20000000	       110 ns/op
BenchmarkSprintfSlowParsingPath-8          	20000000	        73.2 ns/op
BenchmarkSprintfQuoteString-8              	 5000000	       349 ns/op
BenchmarkSprintfInt-8                      	20000000	        73.9 ns/op
BenchmarkSprintfIntInt-8                   	20000000	        99.8 ns/op
BenchmarkSprintfPrefixedInt-8              	10000000	       181 ns/op
BenchmarkSprintfFloat-8                    	10000000	       131 ns/op
BenchmarkSprintfComplex-8                  	 5000000	       399 ns/op
BenchmarkSprintfBoolean-8                  	20000000	        63.6 ns/op
BenchmarkSprintfHexString-8                	 5000000	       332 ns/op
BenchmarkSprintfHexBytes-8                 	 5000000	       420 ns/op
BenchmarkSprintfBytes-8                    	 3000000	       505 ns/op
BenchmarkSprintfStringer-8                 	 5000000	       249 ns/op
BenchmarkSprintfStructure-8                	 2000000	       893 ns/op
BenchmarkManyArgs-8                        	 5000000	       365 ns/op
BenchmarkFprintInt-8                       	10000000	       196 ns/op
BenchmarkFprintfBytes-8                    	 5000000	       254 ns/op
BenchmarkFprintIntNoAlloc-8                	10000000	       199 ns/op
BenchmarkFprintf-8                         	10000000	       171 ns/op
BenchmarkBprintfWithoutPool-8              	20000000	        63.0 ns/op
BenchmarkBprintfWithPool-8                 	20000000	       100 ns/op
BenchmarkSprintf-8                         	20000000	        87.2 ns/op
BenchmarkScanInts-8                        	    3000	    616747 ns/op
BenchmarkScanRecursiveInt-8                	      20	  78370365 ns/op
BenchmarkScanRecursiveIntReaderWrapper-8   	      20	  77659000 ns/op
PASS
ok  	fmt	54.106s
成功: 进程退出代码 0.
