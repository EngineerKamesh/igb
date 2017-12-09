package models

type Product struct {
	SKU                 string
	Name                string
	Description         string
	ThumbnailPreviewURI string
	ImagePreviewURI     string
	Price               float64
	Route               string
	SummaryDetail       string
	Quantity            int
}
