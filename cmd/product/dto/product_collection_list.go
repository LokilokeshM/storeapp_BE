package dto

type ProductCollectionList struct {
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	ID                *int64      `json:"id,omitempty"`
	Title             *string     `json:"title,omitempty"`
	BodyHTML          *string     `json:"body_html,omitempty"`
	Vendor            *string     `json:"vendor,omitempty"`
	ProductType       *string     `json:"product_type,omitempty"`
	CreatedAt         *string     `json:"created_at,omitempty"`
	Handle            *string     `json:"handle,omitempty"`
	UpdatedAt         *string     `json:"updated_at,omitempty"`
	PublishedAt       *string     `json:"published_at,omitempty"`
	TemplateSuffix    interface{} `json:"template_suffix"`
	PublishedScope    *string     `json:"published_scope,omitempty"`
	Tags              *string     `json:"tags,omitempty"`
	AdminGraphqlAPIID *string     `json:"admin_graphql_api_id,omitempty"`
	Options           []Option    `json:"options,omitempty"`
	Images            []Image     `json:"images,omitempty"`
	Image             *Image      `json:"image,omitempty"`
}
