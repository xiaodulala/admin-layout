package models

type Page struct {
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	List      interface{} `json:"list"`
}
