package response

type DetailProductResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type ReviewProductResponse struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	Rating    int `json:"rating"`
}

type ProductResponse struct {
	TotalReview    int     `json:"totalReview"`
	AverageRatings float64 `json:"average_rating"`
	FiveStar       int     `json:"5_star"`
	FourStar       int     `json:"4_star"`
	ThreeStar      int     `json:"3_star"`
	TwoStar        int     `json:"2_star"`
	OneStar        int     `json:"1_star"`
}
