package models

type Link struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

type Response struct {
	OK   bool        `json:"ok"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
