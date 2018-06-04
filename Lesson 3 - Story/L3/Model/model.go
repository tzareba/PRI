package model

type Story struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
}

type StoryOption struct {
	Text    string `json:"text"`
	Chapter string `json:"chapter"`
}
