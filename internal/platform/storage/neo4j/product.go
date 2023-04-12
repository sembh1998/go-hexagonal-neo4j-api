package neo4j

const (
	graphProductNodeLabel = "Product"
)

type graphProduct struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	BarCode string `json:"bar_code"`
	ImgUrl  string `json:"img_url"`
}
