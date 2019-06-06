<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [gobenchmark](#gobenchmark)
- [What it is NOT](#what-it-is-not)
- [Use it](#use-it)
- [Install](#install)
- [Pre-built benchmarks](#pre-built-benchmarks)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# gobenchmark
A collection of benchmarks of basic operation, as a guide for tuning.

When digging deeper and deeper into performance tuning, we found the performance
of some basic operations varies widely, in a counter-intuitive way.
These differences are critical and does impact performance.
Thus we have benchmarked basic operations, to establish a solid guide for people
working on performance critical programs.


# What it is NOT

- Cache missing or memory layout caused performance issues are not included yet.


# Use it

The statistics listed below gives us a intuitive view of basic operation performances.
**But it is strongly recommended to re-benchmark it on your platform before using it**,
since performance may vary on different architectures.


# Install

```sh
go get github.com/openacid/gobenchmark

make ben # run the benchmark
```


# Pre-built benchmarks

```txt
goos: darwin
goarch: amd64
pkg: github.com/openacid/gobenchmark
BenchmarkInt64_Multi_ByConst_Assign-8        	1000000000	         0.41 ns/op
BenchmarkInt64_Multi_ByVar_Assign-8          	1000000000	         0.42 ns/op
BenchmarkInt64_ShiftLeft_ByConst_Assign-8    	1000000000	         0.40 ns/op
BenchmarkInt64_ShiftLeft_ByVar_Assign-8      	1000000000	         0.96 ns/op
BenchmarkInt64_ShiftRight_ByConst_Assign-8   	1000000000	         0.38 ns/op
BenchmarkInt64_ShiftRight_ByVar_Assign-8     	1000000000	         1.51 ns/op
BenchmarkInt64_Div_ByConst_Assign-8          	1000000000	         0.92 ns/op
BenchmarkInt64_Div_ByVar_Assign-8            	100000000	         8.46 ns/op
BenchmarkInt64_Mod_ByConst_Assign-8          	1000000000	         1.12 ns/op
BenchmarkInt64_Mod_ByVar_Assign-8            	100000000	         8.18 ns/op
BenchmarkInt64_Assign-8                      	1000000000	         0.31 ns/op
BenchmarkFloat64_Multi_ByConst_Assign-8      	1000000000	         0.87 ns/op
BenchmarkFloat64_Multi_ByVar_Assign-8        	1000000000	         0.79 ns/op
BenchmarkFloat64_ToInt64_Assign-8            	1000000000	         0.54 ns/op
BenchmarkFloat64_Assign-8                    	1000000000	         0.78 ns/op
PASS
ok  	github.com/openacid/gobenchmark	12.091s
```