package nosedata

type EmptyDatabaseRichTextProp struct {
	RichText struct{} `json:"rich_text"`
}

func (prop EmptyDatabaseRichTextProp) PropType() string {
	return "rich_text"
}

type NoseDatabaseRichTextProp struct {
	RichText []NoseRichText `json:"rich_text"`
}

func (prop NoseDatabaseRichTextProp) PropType() string {
	return "rich_text"
}

//create new db page text column
func NewDBTextColumn(text string) *NoseDatabaseRichTextProp {
	return &NoseDatabaseRichTextProp{
		RichText: NewSingleRichText(text),
	}
}
