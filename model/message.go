package model

type MessageData struct {
	date    string
	name    string
	content string
}

func NewMessageData(date string, name string, content string) *MessageData {
	return &MessageData{date, name, content}
}

func (m *MessageData) GetContent() string {
	return m.content
}

type MessageCountRankResp struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Rank  int    `json:"rank"`
}

func NewMessageCountRankResp(name string, count, rank int) *MessageCountRankResp {
	return &MessageCountRankResp{name, count, rank}
}

type MessageContentCountRankResp struct {
	Name     string   `json:"name"`
	Counts   []int    `json:"count"`
	Contents []string `json:"contents"`
}

func NewMessageContentCountRankResp(name string, count []int, contents []string) *MessageContentCountRankResp {
	return &MessageContentCountRankResp{name, count, contents}
}
