package dto

type ProductCollectsList struct {
	Collects []Collect `json:"collects"`
}

type Collect struct {
	ID           int64       `json:"id"`
	CollectionID int64       `json:"collection_id"`
	ProductID    int64       `json:"product_id"`
	CreatedAt    interface{} `json:"created_at"`
	UpdatedAt    interface{} `json:"updated_at"`
	Position     int64       `json:"position"`
	SortValue    string      `json:"sort_value"`
}
