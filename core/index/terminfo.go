package index

type TermInfo struct {
	Weight    float32
	Df        int32
	MinLength int32
	MaxLength int32
	MaxWeight float32

	MinId int64
	MaxId int64
}

func NewTermInfo(weight float32, df int32,
	minlength int32, maxlength int32,
	maxweight float32,
	minid int64, maxid int64) *TermInfo {
	return &TermInfo{
		Weight:    weight,
		Df:        df,
		MinLength: minlength,
		MaxLength: maxlength,
		MaxWeight: maxweight,
		MinId:     minid,
		MaxId:     maxid,
	}
}

func (info *TermInfo) AddPosting(docNum int64, weight float32, length int32) {
	if info.MinId > docNum {
		info.MinId = docNum
	}
	info.MaxId = docNum
	info.Weight = weight
	info.Df++
	if info.MaxWeight < weight {
		info.MaxWeight = weight
	}
	if length > -1 {
		if info.MinLength > length {
			info.MinLength = length
		}
		if info.MaxLength < length {
			info.MaxLength = length
		}
	}
}
