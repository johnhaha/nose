package nosedata

type TodoBlock struct {
	Object         string `json:"object,omitempty"`
	ID             string `json:"id,omitempty"`
	CreatedTime    string `json:"created_time,omitempty"`
	LastEditedTime string `json:"last_edited_time,omitempty"`
	HasChildren    bool   `json:"has_children,omitempty"`
	Type           string `json:"type"`
	Archived       bool   `json:"archived,omitempty"`
	ToDo           ToDo   `json:"to_do"`
}

type ToDo struct {
	RichText []NoseRichText `json:"rich_text"`
	Checked  bool           `json:"checked"`
}

func NewTodoBlock(text string) *TodoBlock {
	return &TodoBlock{
		Object: "block",
		Type:   "to_do",
		ToDo: ToDo{
			RichText: NewSingleRichText(text),
		},
	}
}

func (block TodoBlock) BlockType() string {
	return "to_do"
}
