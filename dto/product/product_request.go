package productdto

type ProductRequest struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Price      int    `json:"price"`
	Image      string `json:"image"`
	Qty        int    `json:"qty"`
	UserID     int    `json:"user_id"`
	CategoryID int    `json:"category_id"`
}
