package memory

import "testing"

func TestFixLengthStruct(t *testing.T) {
	FixLengthStruct()
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := FixStruct{
			id:    1,
			index: 2,
			size:  3,
		}
		obj.id = obj.size
	}
}
