package example

type Struct struct {
	Value *Value
}

type Value struct {
	value int
}

func New(value int) *Struct {
	return &Struct{
		Value: &Value{
			value: value,
		},
	}
}

func NewSlice() []*Struct {
	return []*Struct{
		New(0),
	}
}
