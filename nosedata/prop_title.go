package nosedata

type NoseDatabaseTitleProp struct {
	Title []NoseRichText `json:"title"`
}

func (prop NoseDatabaseTitleProp) PropType() string {
	return "title"
}

//create new db page title column
func NewDBTitleColumn(text string) *NoseDatabaseTitleProp {
	return &NoseDatabaseTitleProp{
		Title: NewSingleRichText(text),
	}
}

type EmptyDatabaseTitleProp struct {
	Title struct{} `json:"title"`
}

func (prop EmptyDatabaseTitleProp) PropType() string {
	return "title"
}
