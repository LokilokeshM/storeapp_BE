package dto

type AddProductResponse struct {
	ProductResponse *ProductResponse `json:"product,omitempty"`
}

type ProductResponse struct {
	ID                *int64        `json:"id,omitempty"`
	Title             *string       `json:"title,omitempty"`
	BodyHTML          *string       `json:"body_html,omitempty"`
	Vendor            *string       `json:"vendor,omitempty"`
	ProductType       *string       `json:"product_type,omitempty"`
	CreatedAt         *string       `json:"created_at,omitempty"`
	Handle            *string       `json:"handle,omitempty"`
	UpdatedAt         *string       `json:"updated_at,omitempty"`
	PublishedAt       *string       `json:"published_at,omitempty"`
	TemplateSuffix    interface{}   `json:"template_suffix"`
	Status            *string       `json:"status,omitempty"`
	PublishedScope    *string       `json:"published_scope,omitempty"`
	Tags              *string       `json:"tags,omitempty"`
	AdminGraphqlAPIID *string       `json:"admin_graphql_api_id,omitempty"`
	Variants          []Variant     `json:"variants,omitempty"`
	Options           []Option      `json:"options,omitempty"`
	Images            []interface{} `json:"images,omitempty"`
	Image             interface{}   `json:"image"`
}

type Option struct {
	ID        *int64   `json:"id,omitempty"`
	ProductID *int64   `json:"product_id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Position  *int64   `json:"position,omitempty"`
	Values    []string `json:"values,omitempty"`
}

type Variant struct {
	ID                   *int64             `json:"id,omitempty"`
	ProductID            *int64             `json:"product_id,omitempty"`
	Title                *string            `json:"title,omitempty"`
	Price                *string            `json:"price,omitempty"`
	Sku                  *string            `json:"sku,omitempty"`
	Position             *int64             `json:"position,omitempty"`
	InventoryPolicy      *string            `json:"inventory_policy,omitempty"`
	CompareAtPrice       interface{}        `json:"compare_at_price"`
	FulfillmentService   *string            `json:"fulfillment_service,omitempty"`
	InventoryManagement  interface{}        `json:"inventory_management"`
	Option1              *string            `json:"option1,omitempty"`
	Option2              interface{}        `json:"option2"`
	Option3              interface{}        `json:"option3"`
	CreatedAt            *string            `json:"created_at,omitempty"`
	UpdatedAt            *string            `json:"updated_at,omitempty"`
	Taxable              *bool              `json:"taxable,omitempty"`
	Barcode              interface{}        `json:"barcode"`
	Grams                *int64             `json:"grams,omitempty"`
	ImageID              interface{}        `json:"image_id"`
	Weight               *float32           `json:"weight,omitempty"`
	WeightUnit           *string            `json:"weight_unit,omitempty"`
	InventoryItemID      *int64             `json:"inventory_item_id,omitempty"`
	InventoryQuantity    *int64             `json:"inventory_quantity,omitempty"`
	OldInventoryQuantity *int64             `json:"old_inventory_quantity,omitempty"`
	PresentmentPrices    []PresentmentPrice `json:"presentment_prices,omitempty"`
	RequiresShipping     *bool              `json:"requires_shipping,omitempty"`
	AdminGraphqlAPIID    *string            `json:"admin_graphql_api_id,omitempty"`
}

type PresentmentPrice struct {
	Price          *Price      `json:"price,omitempty"`
	CompareAtPrice interface{} `json:"compare_at_price"`
}

type Price struct {
	Amount       *string `json:"amount,omitempty"`
	CurrencyCode *string `json:"currency_code,omitempty"`
}
