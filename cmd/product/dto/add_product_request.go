package dto

type AddProductRequest struct {
	ProductRequest *ProductRequest `json:"product,omitempty"`
}

type ProductRequest struct {
	Title       *string  `json:"title,omitempty"`
	BodyHTML    *string  `json:"body_html,omitempty"`
	Vendor      *string  `json:"vendor,omitempty"`
	ProductType *string  `json:"product_type,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}
