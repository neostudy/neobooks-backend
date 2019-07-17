package models

type Note struct {
	ID     string `json:"id"`
	NType  int    `json:"ntype"`
	PType  int    `json:"ptype"`
	PID    string `json:"pid"`
	BookID string `json:"bookid"`
	ChapID string `json:"chapid"`
	ParaID string `json:"paraid"`
	SentID string `json:"sentid"`
	WordID string `json:"wordid"`
}
