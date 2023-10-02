package brand_domain

type BrandRepo interface {
	GetCriteria(filter *BrandFilter) ([]Criterion, error)
}
