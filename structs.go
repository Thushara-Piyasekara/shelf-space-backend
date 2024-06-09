package main

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

type author struct {
	ID     string  `json:"id"`
	Name   string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
