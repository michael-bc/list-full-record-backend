package brand_http

type CriterionDTO struct {
	Category CategoryDTO `json:"category"`
	Location LocationDTO `json:"location"`
	Brands   []BrandDTO  `json:"brands"`
}
