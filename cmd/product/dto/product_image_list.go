package dto

type ProductImageLIst struct {
	Images []Image `json:"images,omitempty"`
}

type Image struct {
	ID                *int64        `json:"id,omitempty"`
	ProductID         *int64        `json:"product_id,omitempty"`
	Position          *int64        `json:"position,omitempty"`
	CreatedAt         *string       `json:"created_at,omitempty"`
	UpdatedAt         *string       `json:"updated_at,omitempty"`
	Alt               interface{}   `json:"alt"`
	Width             *int64        `json:"width,omitempty"`
	Height            *int64        `json:"height,omitempty"`
	Src               *string       `json:"src,omitempty"`
	VariantIDS        []interface{} `json:"variant_ids,omitempty"`
	AdminGraphqlAPIID *string       `json:"admin_graphql_api_id,omitempty"`
}
