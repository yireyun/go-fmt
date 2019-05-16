c:/go/bin/go.exe test -test.bench=.* [C:/Go/src/fmt]
go1.7.6
BenchmarkSprintfPadding-8                  	10000000	       230 ns/op
BenchmarkSprintfEmpty-8                    	50000000	        29.0 ns/op
BenchmarkSprintfString-8                   	20000000	        82.2 ns/op
BenchmarkSprintfTruncateString-8           	20000000	        97.6 ns/op
BenchmarkSprintfQuoteString-8              	10000000	       329 ns/op
BenchmarkSprintfInt-8                      	20000000	       104 ns/op
BenchmarkSprintfIntInt-8                   	10000000	       148 ns/op
BenchmarkSprintfPrefixedInt-8              	10000000	       249 ns/op
BenchmarkSprintfFloat-8                    	10000000	       168 ns/op
BenchmarkSprintfComplex-8                  	 3000000	       440 ns/op
BenchmarkSprintfBoolean-8                  	20000000	        90.1 ns/op
BenchmarkSprintfHexString-8                	 3000000	       388 ns/op
BenchmarkSprintfHexBytes-8                 	 3000000	       396 ns/op
BenchmarkSprintfBytes-8                    	 3000000	       557 ns/op
BenchmarkSprintfStringer-8                 	 5000000	       333 ns/op
BenchmarkSprintfStructure-8                	 2000000	      1029 ns/op
BenchmarkManyArgs-8                        	 3000000	       592 ns/op
BenchmarkFprintInt-8                       	 5000000	       272 ns/op
BenchmarkFprintfBytes-8                    	 5000000	       344 ns/op
BenchmarkFprintIntNoAlloc-8                	 5000000	       250 ns/op
BenchmarkFprintf-8                         	 5000000	       282 ns/op
BenchmarkBprintfWithoutPool-8              	10000000	       150 ns/op
BenchmarkBprintfWithPool-8                 	10000000	       188 ns/op
BenchmarkSprintf-8                         	20000000	       140 ns/op
BenchmarkScanInts-8                        	    2000	    933027 ns/op
BenchmarkScanRecursiveInt-8                	      20	  78260850 ns/op
BenchmarkScanRecursiveIntReaderWrapper-8   	      20	  81296695 ns/op
PASS
ok  	fmt	55.845s
成功: 进程退出代码 0.
