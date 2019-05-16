c:/go/bin/go.exe test -test.bench=.* [C:/Go/src/fmt]
go1.9.7
goos: windows
goarch: amd64
pkg: fmt
BenchmarkSprintfPadding-8                  	20000000	       143 ns/op
BenchmarkSprintfEmpty-8                    	200000000	        17.7 ns/op
BenchmarkSprintfString-8                   	20000000	        61.9 ns/op
BenchmarkSprintfTruncateString-8           	20000000	       106 ns/op
BenchmarkSprintfQuoteString-8              	 5000000	       313 ns/op
BenchmarkSprintfInt-8                      	20000000	        68.7 ns/op
BenchmarkSprintfIntInt-8                   	20000000	        97.1 ns/op
BenchmarkSprintfPrefixedInt-8              	10000000	       165 ns/op
BenchmarkSprintfFloat-8                    	10000000	       132 ns/op
BenchmarkSprintfComplex-8                  	 5000000	       376 ns/op
BenchmarkSprintfBoolean-8                  	20000000	        62.0 ns/op
BenchmarkSprintfHexString-8                	 5000000	       344 ns/op
BenchmarkSprintfHexBytes-8                 	 5000000	       393 ns/op
BenchmarkSprintfBytes-8                    	 3000000	       514 ns/op
BenchmarkSprintfStringer-8                 	 5000000	       250 ns/op
BenchmarkSprintfStructure-8                	 2000000	       849 ns/op
BenchmarkManyArgs-8                        	 5000000	       365 ns/op
BenchmarkFprintInt-8                       	10000000	       198 ns/op
BenchmarkFprintfBytes-8                    	 5000000	       250 ns/op
BenchmarkFprintIntNoAlloc-8                	10000000	       206 ns/op
BenchmarkFprintf-8                         	10000000	       174 ns/op
BenchmarkBprintfWithoutPool-8              	20000000	        59.1 ns/op
BenchmarkBprintfWithPool-8                 	20000000	        94.0 ns/op
BenchmarkSprintf-8                         	20000000	        88.0 ns/op
BenchmarkScanInts-8                        	    2000	    640093 ns/op
BenchmarkScanRecursiveInt-8                	      20	  71003070 ns/op
BenchmarkScanRecursiveIntReaderWrapper-8   	      20	  74852990 ns/op
PASS
ok  	fmt	53.696s
成功: 进程退出代码 0.
