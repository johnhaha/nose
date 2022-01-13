package nosedata

type NoseNumber struct {
	Number interface{} `json:"number"`
}

func (n *NoseNumber) PropType() string {
	return "number"
}

//for database prop init
type EmptyDatabaseNumberProp struct {
	Number struct{} `json:"number"`
}

func (n EmptyDatabaseNumberProp) PropType() string {
	return "number"
}
