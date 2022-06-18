package nosedata

type BlockInfo struct {
	Object         string `json:"object"`
	ID             string `json:"id"`
	CreatedTime    string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
	HasChildren    bool   `json:"has_children"`
	Type           string `json:"type"`
	Archived       bool   `json:"archived"`
}

type ParagraphBlock struct {
	Object         string    `json:"object,omitempty"`
	ID             string    `json:"id,omitempty"`
	CreatedTime    string    `json:"created_time,omitempty"`
	LastEditedTime string    `json:"last_edited_time,omitempty"`
	HasChildren    bool      `json:"has_children,omitempty"`
	Type           string    `json:"type"`
	Archived       bool      `json:"archived,omitempty"`
	Paragraph      Paragraph `json:"paragraph"`
}

type Paragraph struct {
	RichText []NoseRichText `json:"rich_text"`
}

func (block ParagraphBlock) BlockType() string {
	return "paragraph"
}

func NewParagraphBlock(texts ...string) ParagraphBlock {
	richTexts := make([]NoseRichText, len(texts))
	for i, text := range texts {
		richTexts[i] = *NewRichText(text)
	}
	return ParagraphBlock{
		Object:    "block",
		Type:      "paragraph",
		Paragraph: Paragraph{RichText: richTexts},
	}
}

func NewParagraphDataBlock(data any) ParagraphBlock {
	richTexts := NewRichTextFromData(data)
	return ParagraphBlock{
		Object:    "block",
		Type:      "paragraph",
		Paragraph: Paragraph{RichText: richTexts},
	}
}
