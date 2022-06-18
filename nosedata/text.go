package nosedata

import (
	"log"
	"reflect"
	"strings"
)

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color,omitempty"`
}

func (a *Annotations) FindColor(tags []string) *Annotations {
	for _, t := range tags {
		if isNoseColor(t) {
			a.Color = t
			return a
		}
	}
	return a
}

func (a *Annotations) FindBold(tags []string) *Annotations {
	for _, t := range tags {
		if t == boldTag {
			a.Bold = true
			return a
		}
	}
	return a
}

func (a *Annotations) FindCode(tags []string) *Annotations {
	for _, t := range tags {
		if t == textCodeTag {
			a.Code = true
			return a
		}
	}
	return a
}

type NoseText struct {
	Content string `json:"content,omitempty"`
}

type NoseRichText struct {
	Type        string       `json:"type,omitempty"`
	Text        NoseText     `json:"text,omitempty"`
	PlainText   string       `json:"plain_text,omitempty"`
	Href        string       `json:"href,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
}

func NewRichText(text string) *NoseRichText {
	return &NoseRichText{Text: NoseText{Content: text}}
}

type TitleText struct {
	Title []NoseRichText `json:"title"`
}

func NewSingleRichText(text string) []NoseRichText {
	if text == "" {
		return nil
	}
	return []NoseRichText{
		{Text: NoseText{Content: text}},
	}
}

func NewRichTextFromData(data any) []NoseRichText {
	var richText []NoseRichText
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).IsZero() {
			continue
		}
		f := t.Field(i)
		switch typedValue := v.Field(i).Interface().(type) {
		case string:
			t := NoseRichText{Text: NoseText{Content: typedValue}}
			if n, ok := f.Tag.Lookup("nose"); ok {
				tagList := strings.Split(n, ",")
				a := new(Annotations)
				a.FindColor(tagList).FindBold(tagList).FindCode(tagList)
				t.Annotations = a
			}
			richText = append(richText, t)
		default:
			log.Printf("unsupported field %v", f.Name)
		}
	}
	return richText
}
