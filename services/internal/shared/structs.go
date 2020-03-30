package shared

// TikaResponse represents a response from tikad
type TikaResponse struct {
	Body         string `json:"body"`
	DocumentType string `json:"documentType"`
}
