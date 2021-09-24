package nosedata

type DBInfo struct {
	PageType   string `json:"-"`
	DatabaseID string `json:"database_id"`
}

type EmptyDatabaseTitleProp struct {
	Title struct{} `json:"title"`
}

func (prop EmptyDatabaseTitleProp) PropType() string {
	return "title"
}

type EmptyDatabaseRichTextProp struct {
	RichText struct{} `json:"rich_text"`
}

func (prop EmptyDatabaseRichTextProp) PropType() string {
	return "rich_text"
}

type NoseDatabaseTitleProp struct {
	Title []NoseRichText `json:"title"`
}

func (prop NoseDatabaseTitleProp) PropType() string {
	return "title"
}

type NoseDatabaseRichTextProp struct {
	RichText []NoseRichText `json:"rich_text"`
}

func (prop NoseDatabaseRichTextProp) PropType() string {
	return "rich_text"
}

//create new db page title column
func NewDBTitleColumn(text string) *NoseDatabaseTitleProp {
	return &NoseDatabaseTitleProp{
		Title: NewSingleRichText(text),
	}
}

//create new db page text column
func NewDBTextColumn(text string) *NoseDatabaseRichTextProp {
	return &NoseDatabaseRichTextProp{
		RichText: NewSingleRichText(text),
	}
}
