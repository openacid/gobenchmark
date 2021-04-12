<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [gobenchmark](#gobenchmark)
- [What it is NOT](#what-it-is-not)
- [Use it](#use-it)
- [Install](#install)
- [Tips for tuning](#tips-for-tuning)
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

# Tips for tuning

According to the benchmark, there are some general tips for tuning computation
dense programs:

-   Usually using static data is faster:
    -   **Most of the time**, operating on a const is faster than on a variable.
    -   Using array is (about 2 times) faster than using a slice.

-   Integer are faster than float number.

-   Usually `/` and `%` are very slow. Use shift if possible.



# Pre-built benchmarks

```txt
goos: darwin 
goarch: amd64 
pkg: github.com/openacid/gobenchmark 
BenchmarkInt64_And_ByConst_Assign-8                   0.364 ns/op -
BenchmarkInt64_And_ByVar_Assign-8                     0.397 ns/op -
BenchmarkInt64_Or_ByConst_Assign-8                    0.363 ns/op -
BenchmarkInt64_Or_ByVar_Assign-8                      0.400 ns/op --
BenchmarkInt64_Xor_ByConst_Assign-8                   0.367 ns/op -
BenchmarkInt64_Xor_ByVar_Assign-8                     0.409 ns/op --
BenchmarkInt64_Add_ByConst_Assign-8                   0.528 ns/op --
BenchmarkInt64_Add_ByVar_Assign-8                     0.394 ns/op -
BenchmarkInt64_Sub_ByConst_Assign-8                   0.526 ns/op --
BenchmarkInt64_Sub_ByVar_Assign-8                     0.400 ns/op --
BenchmarkInt64_Multi_ByConst_Assign-8                 0.404 ns/op --
BenchmarkInt64_Multi_ByVar_Assign-8                   0.428 ns/op --
BenchmarkInt64_ShiftLeft_ByConst_Assign-8             0.359 ns/op -
BenchmarkInt64_ShiftLeft_ByVar_Assign-8               0.952 ns/op ----
BenchmarkInt64_ShiftRight_ByConst_Assign-8            0.362 ns/op -
BenchmarkInt64_ShiftRight_ByVar_Assign-8              1.45 ns/op -------
BenchmarkInt64_Div_ByConst_Assign-8                   1.01 ns/op -----
BenchmarkInt64_Div_ByVar_Assign-8                     8.18 ns/op ----------------------------------------
BenchmarkInt64_Mod_ByConst_Assign-8                   1.05 ns/op -----
BenchmarkInt64_Mod_ByVar_Assign-8                     8.46 ns/op ------------------------------------------
BenchmarkInt64_Assign-8                               0.268 ns/op -
BenchmarkFloat64_Multi_ByConst_Assign-8               0.811 ns/op ----
BenchmarkFloat64_Multi_ByVar_Assign-8                 0.813 ns/op ----
BenchmarkFloat64_Div_ByConst_Assign-8                 3.73 ns/op ------------------
BenchmarkFloat64_Div_ByVar_Assign-8                   3.68 ns/op ------------------
BenchmarkFloat64_ToInt64_Assign-8                     0.554 ns/op --
BenchmarkFloat64_Assign-8                             0.777 ns/op ---
BenchmarkInt64Array_Get_Assign-8                      0.265 ns/op -
BenchmarkInt64Slice_Get_Assign-8                      0.535 ns/op --
PASS 
ok  	github.com/openacid/gobenchmark.221s 
```