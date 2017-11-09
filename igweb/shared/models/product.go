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

type Products []*Product

func (p Products) Len() int           { return len(p) }
func (p Products) Less(i, j int) bool { return p[i].Price < p[j].Price }
func (p Products) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
