package models

type Ofac struct {
	Name    string `json:"name"`
	Link    string `json:"link"`
	List    string `json:"list"`
	Program string `json:"program"`
	Score   string `json:"score"`
}
