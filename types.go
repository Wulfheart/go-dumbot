package hadrian

type Weighting int64

type Parameters struct {
	PowerSize_A int
	PowerSize_B int
	PowerSize_C int
}


func StandardParameters() Parameters {
	return Parameters{
		PowerSize_A: 1,
		PowerSize_B: 4,
		PowerSize_C: 16,
	}
}
