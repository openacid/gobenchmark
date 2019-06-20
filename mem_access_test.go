package gobenchmark

import "testing"

const nelt = 1024 * 1024
const neltMask = nelt - 1

var InpI64Arr = [nelt]int64{}
var InpI64Slice = InpI64Arr[0:nelt]

func BenchmarkInt64Array_Get_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s = InpI64Arr[5]
	}
	Output = int(s)
}

func BenchmarkInt64Slice_Get_Assign(b *testing.B) {
	var s int64
	for i := 0; i < b.N; i++ {
		s = InpI64Slice[5]
	}
	Output = int(s)
}
