c:/go/bin/go.exe test -test.bench=.* [C:/Go/src/fmt]
go1.11.10
goos: windows
goarch: amd64
pkg: fmt
BenchmarkSprintfPadding-8                  	10000000	       161 ns/op
BenchmarkSprintfEmpty-8                    	100000000	        16.8 ns/op
BenchmarkSprintfString-8                   	30000000	        48.4 ns/op
BenchmarkSprintfTruncateString-8           	20000000	        86.3 ns/op
BenchmarkSprintfSlowParsingPath-8          	30000000	        56.3 ns/op
BenchmarkSprintfQuoteString-8              	 5000000	       311 ns/op
BenchmarkSprintfInt-8                      	50000000	        34.7 ns/op
BenchmarkSprintfIntInt-8                   	20000000	        71.0 ns/op
BenchmarkSprintfPrefixedInt-8              	10000000	       165 ns/op
BenchmarkSprintfFloat-8                    	10000000	       116 ns/op
BenchmarkSprintfComplex-8                  	 5000000	       311 ns/op
BenchmarkSprintfBoolean-8                  	30000000	        43.8 ns/op
BenchmarkSprintfHexString-8                	20000000	       160 ns/op
BenchmarkSprintfHexBytes-8                 	10000000	       180 ns/op
BenchmarkSprintfBytes-8                    	 5000000	       305 ns/op
BenchmarkSprintfStringer-8                 	10000000	       220 ns/op
BenchmarkSprintfStructure-8                	 2000000	       688 ns/op
BenchmarkManyArgs-8                        	 5000000	       232 ns/op
BenchmarkFprintInt-8                       	10000000	       155 ns/op
BenchmarkFprintfBytes-8                    	 5000000	       236 ns/op
BenchmarkFprintIntNoAlloc-8                	10000000	       162 ns/op
BenchmarkFprintf-8                         	10000000	       158 ns/op
BenchmarkBprintfWithoutPool-8              	20000000	        87.9 ns/op
BenchmarkBprintfWithPool-8                 	10000000	       122 ns/op
BenchmarkSprintf-8                         	30000000	        61.8 ns/op
BenchmarkScanInts-8                        	    3000	    615046 ns/op
BenchmarkScanRecursiveInt-8                	      20	  69664755 ns/op
BenchmarkScanRecursiveIntReaderWrapper-8   	      20	  79616970 ns/op
PASS
ok  	fmt	54.294s
成功: 进程退出代码 0.
