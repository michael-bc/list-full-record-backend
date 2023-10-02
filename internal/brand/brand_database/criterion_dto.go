package brand_database

type criterionDTO struct {
	Category categoryDTO `json:"category"`
	Location locationDTO `json:"location"`
	Brands   brandDTOs   `json:"brands"`
}
