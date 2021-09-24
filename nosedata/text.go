package nosedata

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color,omitempty"`
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
