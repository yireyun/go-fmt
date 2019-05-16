c:/go/bin/go.exe test -test.bench=.* [C:/Go/src/fmt]
go1.8.7
BenchmarkSprintfPadding-8                  	10000000	       214 ns/op
BenchmarkSprintfEmpty-8                    	50000000	        27.3 ns/op
BenchmarkSprintfString-8                   	20000000	       103 ns/op
BenchmarkSprintfTruncateString-8           	10000000	       145 ns/op
BenchmarkSprintfQuoteString-8              	 5000000	       335 ns/op
BenchmarkSprintfInt-8                      	20000000	        94.4 ns/op
BenchmarkSprintfIntInt-8                   	10000000	       133 ns/op
BenchmarkSprintfPrefixedInt-8              	10000000	       207 ns/op
BenchmarkSprintfFloat-8                    	10000000	       159 ns/op
BenchmarkSprintfComplex-8                  	 5000000	       423 ns/op
BenchmarkSprintfBoolean-8                  	20000000	        79.7 ns/op
BenchmarkSprintfHexString-8                	 5000000	       332 ns/op
BenchmarkSprintfHexBytes-8                 	 3000000	       352 ns/op
BenchmarkSprintfBytes-8                    	 3000000	       515 ns/op
BenchmarkSprintfStringer-8                 	 5000000	       282 ns/op
BenchmarkSprintfStructure-8                	 2000000	       940 ns/op
BenchmarkManyArgs-8                        	 3000000	       543 ns/op
BenchmarkFprintInt-8                       	 5000000	       251 ns/op
BenchmarkFprintfBytes-8                    	 5000000	       299 ns/op
BenchmarkFprintIntNoAlloc-8                	10000000	       236 ns/op
BenchmarkFprintf-8                         	 5000000	       254 ns/op
BenchmarkBprintfWithoutPool-8              	10000000	       134 ns/op
BenchmarkBprintfWithPool-8                 	10000000	       164 ns/op
BenchmarkSprintf-8                         	20000000	       117 ns/op
BenchmarkScanInts-8                        	    2000	    637510 ns/op
BenchmarkScanRecursiveInt-8                	      20	  78799210 ns/op
BenchmarkScanRecursiveIntReaderWrapper-8   	      20	  79006150 ns/op
PASS
ok  	fmt	52.722s
成功: 进程退出代码 0.
