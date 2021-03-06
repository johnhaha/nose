package nosedata

import (
	"log"
	"reflect"
	"time"
)

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

type CreateDatabasePageReq struct {
	Parent     DBInfo                      `json:"parent"`
	Title      []NoseRichText              `json:"title,omitempty"`
	Properties map[string]NoseDatabaseProp `json:"properties"`
}

type AppendBlockReq struct {
	Children []NoseBlock `json:"children"`
}

func NewAppendTextBlockReq(texts ...string) *AppendBlockReq {
	return &AppendBlockReq{
		Children: []NoseBlock{NewParagraphBlock(texts...)},
	}
}

//new append text block from struct data
func NewAppendDataTextBlockReq(data any) *AppendBlockReq {
	return &AppendBlockReq{
		Children: []NoseBlock{NewParagraphDataBlock(data)},
	}
}

func NewAppendTodoBlockReq(text string) *AppendBlockReq {
	return &AppendBlockReq{
		Children: []NoseBlock{NewTodoBlock(text)},
	}
}

//generate new db page req with struct value
func NewEmptyDatabaseReqX(parentID string, name string, data any) *CreateDatabaseReq {
	pageParent := PageInfo{
		PageType: "page_id",
		PageID:   parentID,
	}
	title := []NoseRichText{
		{Text: NoseText{Content: name}},
	}
	props := make(map[string]NoseDatabaseProp)
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if n, ok := f.Tag.Lookup("nose"); ok && n == "title" {
			props[f.Name] = EmptyDatabaseTitleProp{}
			continue
		}
		value := v.Field(i).Interface()
		switch value.(type) {
		case string:
			props[f.Name] = EmptyDatabaseRichTextProp{}
		case int, int8, int16, int32, int64, float32, float64:
			props[f.Name] = EmptyDatabaseNumberProp{}
		case time.Time:
			props[f.Name] = EmptyDatabaseDateProp{}
		}
	}
	req := CreateDatabaseReq{
		Parent:     pageParent,
		Title:      title,
		Properties: props,
	}
	return &req
}

//generate db page req with struct
func NewCreateDatabasePageReqX(dbID string, data any) *CreateDatabasePageReq {
	parent := DBInfo{
		DatabaseID: dbID,
	}
	props := make(map[string]NoseDatabaseProp)

	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if v.Field(i).IsZero() {
			continue
		}
		value := v.Field(i).Interface()
		if n, ok := f.Tag.Lookup("nose"); ok && n == "title" {
			switch typedValue := value.(type) {
			case string:
				props[f.Name] = NewDBTitleColumn(typedValue)
			default:
				props[f.Name] = NewDBTitleColumn("unsupported field value type")
				log.Printf("unsupported field value for field %v, value %v", f.Name, typedValue)
			}
		} else {
			switch typedValue := value.(type) {
			case string:
				props[f.Name] = NewDBTextColumn(typedValue)
			case int, int8, int16, int32, int64, float32, float64:
				props[f.Name] = NewDBNumberColumn(typedValue)
			case time.Time:
				props[f.Name] = NewDBDateColumn(typedValue)
			default:
				log.Printf("unsupported field value for field %v, value %v", f.Name, typedValue)
			}
		}
	}

	return &CreateDatabasePageReq{
		Parent:     parent,
		Properties: props,
	}
}

//generate new data base page req
// func NewCreateDatabasePageReq(dbID string, title map[string]string, values map[string]any) *CreateDatabasePageReq {
// 	parent := DBInfo{
// 		DatabaseID: dbID,
// 	}
// 	props := make(map[string]NoseDatabaseProp)
// 	for key, value := range title {
// 		props[key] = NewDBTitleColumn(value)
// 	}
// 	for key, value := range values {
// 		switch v := value.(type) {
// 		case string:
// 			props[key] = NewDBTextColumn(v)
// 		case int, int16, int32, int64, int8, float32, float64:
// 			props[key] = NewDBNumberColumn(v)
// 		}
// 	}
// 	return &CreateDatabasePageReq{
// 		Parent:     parent,
// 		Properties: props,
// 	}
// }

// func NewEmptyDatabaseReq(parentID string, name string, columnTitle string, textColumns ...string) *CreateDatabaseReq {
// 	pageParent := PageInfo{
// 		PageType: "page_id",
// 		PageID:   parentID,
// 	}
// 	title := []NoseRichText{
// 		{Text: NoseText{Content: name}},
// 	}
// 	props := make(map[string]NoseDatabaseProp)
// 	for _, text := range textColumns {
// 		props[text] = EmptyDatabaseRichTextProp{}
// 	}
// 	props[columnTitle] = EmptyDatabaseTitleProp{}
// 	req := CreateDatabaseReq{
// 		Parent:     pageParent,
// 		Title:      title,
// 		Properties: props,
// 	}
// 	return &req
// }
