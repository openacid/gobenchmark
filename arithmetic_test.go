package gobenchmark

import "testing"

// Exported (global) variable to store function results
// during benchmarking to ensure side-effect free calls
// are not optimized away.
var Output int

var Inp1I64 int64 = 3
var Inp2I64 int64 = 5
var Inp1U uint = 5
var Inp1Float64 float64 = 1.23456789
var Inp2Float64 float64 = 1.23456789

func BenchmarkInt64_Multi_ByConst_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += Inp1I64 * 87
	}
	Output = int(s)
}

func BenchmarkInt64_Multi_ByVar_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += Inp1I64 * Inp2I64
	}
	Output = int(s)
}

func BenchmarkInt64_ShiftLeft_ByConst_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += Inp1I64 << 5
	}
	Output = int(s)
}

func BenchmarkInt64_ShiftLeft_ByVar_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += Inp1I64 << Inp1U
	}
	Output = int(s)
}

func BenchmarkInt64_ShiftRight_ByConst_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += Inp1I64 >> 5
	}
	Output = int(s)
}

func BenchmarkInt64_ShiftRight_ByVar_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += Inp1I64 >> Inp1U
	}
	Output = int(s)
}

func BenchmarkInt64_Div_ByConst_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += Inp1I64 / 3
	}
	Output = int(s)
}

func BenchmarkInt64_Div_ByVar_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += Inp1I64 / Inp2I64
	}
	Output = int(s)
}

func BenchmarkInt64_Mod_ByConst_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += Inp1I64 % 3
	}
	Output = int(s)
}

func BenchmarkInt64_Mod_ByVar_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += Inp1I64 % Inp2I64
	}
	Output = int(s)
}

func BenchmarkInt64_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += Inp1I64
	}
	Output = int(s)
}

func BenchmarkFloat64_Multi_Assign(b *testing.B) {
	var s float64
	for i := 0; i < b.N; i++ {
		s += Inp1Float64 * Inp2Float64
	}
	Output = int(s)
}

func BenchmarkFloat64_ToInt64_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s += int64(Inp1Float64 * Inp2Float64)
	}
	Output = int(s)
}

func BenchmarkFloat64_Assign(b *testing.B) {
	var s float64
	for i := 0; i < b.N; i++ {
		s += Inp1Float64
	}
	Output = int(s)
}
