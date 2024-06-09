package constants

type book struct {
	ID     int64  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

type author struct {
	ID        int64 `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
}
