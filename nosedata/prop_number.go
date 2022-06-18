package nosedata

type NoseNumber struct {
	Number any `json:"number"`
}

func (n *NoseNumber) PropType() string {
	return "number"
}

//create new db page number column
func NewDBNumberColumn(n any) *NoseNumber {
	return &NoseNumber{Number: n}
}

//for database prop init
type EmptyDatabaseNumberProp struct {
	Number struct{} `json:"number"`
}

func (n EmptyDatabaseNumberProp) PropType() string {
	return "number"
}
