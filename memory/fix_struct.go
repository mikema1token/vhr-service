package memory

type FixStruct struct {
	id    uint32
	index uint32
	size  uint32
}

func FixLengthStruct() {
	_ = FixStruct{
		id:    1,
		index: 2,
		size:  3,
	}
}
