package dto

type Product struct {
	ID int `json:"id,omitempty"`
	// Series holds the value of the "series" field.
	Series string `json:"series,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// Sku holds the value of the "sku" field.
	Sku string `json:"sku,omitempty"`
	// Properties holds the value of the "properties" field.
	Properties map[string]string `json:"properties,omitempty"`
}
