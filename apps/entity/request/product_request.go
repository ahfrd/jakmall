package request

type DetailProductEntity struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type ReviewProductEntity struct {
	Id        string `json:"id"`
	ProductId string `json:"product_id"`
	Rating    string `json:"rating"`
}

type ProductRequest struct {
	DetailProduct []DetailProductEntity `json:"detailProduct"`
	ReviewProduct []ReviewProductEntity `json:"reviewProduct"`
}
