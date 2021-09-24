package nosedata

type NoseDatabaseProp interface {
	PropType() string
}

type NoseBlock interface {
	BlockType() string
}
