package nosedata

type CreatePageReq struct {
	Parent     PageInfo  `json:"parent"`
	Properties TitleText `json:"properties"`
}

func NewEmptyPageReq(parentID string, name string) *CreatePageReq {
	pageParent := PageInfo{
		PageType: "page_id",
		PageID:   parentID,
	}
	title := TitleText{
		Title: []NoseRichText{
			{Text: NoseText{Content: name}},
		},
	}
	req := CreatePageReq{Parent: pageParent, Properties: title}
	return &req
}

type CreateDatabaseReq struct {
	Parent     PageInfo                    `json:"parent"`
	Title      []NoseRichText              `json:"title"`
	Properties map[string]NoseDatabaseProp `json:"properties"`
}

func NewEmptyDatabaseReq(parentID string, name string, columnTitle string, textColumns ...string) *CreateDatabaseReq {
	pageParent := PageInfo{
		PageType: "page_id",
		PageID:   parentID,
	}
	title := []NoseRichText{
		{Text: NoseText{Content: name}},
	}
	props := make(map[string]NoseDatabaseProp)
	for _, text := range textColumns {
		props[text] = EmptyDatabaseRichTextProp{}
	}
	props[columnTitle] = EmptyDatabaseTitleProp{}
	req := CreateDatabaseReq{
		Parent:     pageParent,
		Title:      title,
		Properties: props,
	}
	return &req
}

type CreateDatabasePageReq struct {
	Parent     DBInfo                      `json:"parent"`
	Title      []NoseRichText              `json:"title,omitempty"`
	Properties map[string]NoseDatabaseProp `json:"properties"`
}

func NewCreateDatabasePageReq(dbID string, title map[string]string, values map[string]string) *CreateDatabasePageReq {
	parent := DBInfo{
		DatabaseID: dbID,
	}
	props := make(map[string]NoseDatabaseProp)
	for key, value := range title {
		props[key] = NewDBTitleColumn(value)
	}
	for key, value := range values {
		props[key] = NewDBTextColumn(value)
	}
	return &CreateDatabasePageReq{
		Parent:     parent,
		Properties: props,
	}
}

type AppendBlockReq struct {
	Children []NoseBlock `json:"children"`
}

func NewAppendTextBlockReq(texts ...string) *AppendBlockReq {
	return &AppendBlockReq{
		Children: []NoseBlock{NewParagraphBlock(texts...)},
	}
}

func NewAppendTodoBlockReq(text string) *AppendBlockReq {
	return &AppendBlockReq{
		Children: []NoseBlock{NewTodoBlock(text)},
	}
}
