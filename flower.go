package main

//Content model
type Content struct {
	Page Page     `json:"page"`
	Data []Flower `json:"data"`
}

//Page model
type Page struct {
	Index     int32 `json:"index"`
	PageSize  int32 `json:"pageSize"`
	Total     int32 `json:"total"`
	PageTotal int32 `json:"pageTotal"`
}

//Flower Model
type Flower struct {
	Map     string `json:"map"`
	Unit    string `json:"unit"`
	Price   string `json:"price"`
	Time    int64  `json:"time"`
	Created string `json:"created_str"`
}
