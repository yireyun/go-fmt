c:/go/bin/go.exe test -test.bench=.* [C:/Go/src/fmt]
go1.11.10
goos: windows
goarch: amd64
pkg: fmt
BenchmarkSprintfPadding-8                  	20000000	        99.5 ns/op
BenchmarkSprintfEmpty-8                    	100000000	        18.1 ns/op
BenchmarkSprintfString-8                   	30000000	        53.0 ns/op
BenchmarkSprintfTruncateString-8           	20000000	        96.3 ns/op
BenchmarkSprintfSlowParsingPath-8          	20000000	        62.3 ns/op
BenchmarkSprintfQuoteString-8              	 5000000	       319 ns/op
BenchmarkSprintfInt-8                      	50000000	        33.8 ns/op
BenchmarkSprintfIntInt-8                   	20000000	        71.3 ns/op
BenchmarkSprintfPrefixedInt-8              	10000000	       180 ns/op
BenchmarkSprintfFloat-8                    	20000000	       118 ns/op
BenchmarkSprintfComplex-8                  	 5000000	       306 ns/op
BenchmarkSprintfBoolean-8                  	30000000	        43.9 ns/op
BenchmarkSprintfHexString-8                	20000000	       144 ns/op
BenchmarkSprintfHexBytes-8                 	10000000	       205 ns/op
BenchmarkSprintfBytes-8                    	 5000000	       304 ns/op
BenchmarkSprintfStringer-8                 	10000000	       226 ns/op
BenchmarkSprintfStructure-8                	 2000000	       770 ns/op
BenchmarkManyArgs-8                        	 5000000	       249 ns/op
BenchmarkFprintInt-8                       	10000000	       158 ns/op
BenchmarkFprintfBytes-8                    	 5000000	       242 ns/op
BenchmarkFprintIntNoAlloc-8                	10000000	       168 ns/op
BenchmarkFprint-8                          	10000000	       152 ns/op
BenchmarkFmtPoolFprint-8                   	10000000	       137 ns/op
BenchmarkFmtFprint-8                       	20000000	       100 ns/op
BenchmarkFprintln-8                        	10000000	       151 ns/op
BenchmarkFmtPoolFprintLn-8                 	10000000	       133 ns/op
BenchmarkFmtFprintLn-8                     	20000000	        98.5 ns/op
BenchmarkFprintf-8                         	10000000	       152 ns/op
BenchmarkFmtPoolFprintf-8                  	10000000	       133 ns/op
BenchmarkFmtFprintf-8                      	20000000	       101 ns/op
BenchmarkBprint-8                          	10000000	       140 ns/op
BenchmarkFmtPoolBprint-8                   	20000000	       124 ns/op
BenchmarkFmtBprint-8                       	20000000	        85.5 ns/op
BenchmarkBprintln-8                        	10000000	       140 ns/op
BenchmarkFmtPoolBprintLn-8                 	20000000	       117 ns/op
BenchmarkFmtBprintLn-8                     	20000000	        81.5 ns/op
BenchmarkBprintf-8                         	10000000	       144 ns/op
BenchmarkFmtPoolBprintf-8                  	10000000	       128 ns/op
BenchmarkFmtBprintf-8                      	20000000	        94.7 ns/op
BenchmarkSprint-8                          	10000000	       180 ns/op
BenchmarkFmtPoolSprint-8                   	10000000	       167 ns/op
BenchmarkFmtSprint-8                       	10000000	       126 ns/op
BenchmarkSprintln-8                        	10000000	       172 ns/op
BenchmarkFmtPoolSprintLn-8                 	10000000	       158 ns/op
BenchmarkFmtSprintLn-8                     	10000000	       124 ns/op
BenchmarkSprintf-8                         	30000000	        65.8 ns/op
BenchmarkFmtPoolSprintf-8                  	10000000	       164 ns/op
BenchmarkFmtSprintf-8                      	10000000	       124 ns/op
BenchmarkScanInts-8                        	    3000	    571346 ns/op
BenchmarkScanRecursiveInt-8                	      20	  74517100 ns/op
BenchmarkScanRecursiveIntReaderWrapper-8   	      20	  70264810 ns/op
PASS
ok  	fmt	93.547s
成功: 进程退出代码 0.
