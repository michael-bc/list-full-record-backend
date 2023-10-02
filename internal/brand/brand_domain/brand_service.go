package brand_domain

type BrandService interface {
	GetCriteria(filter *BrandFilter) ([]Criterion, error)
}

type BrandSvcImpl struct {
	brandRepo BrandRepo
}

func NewBrandServiceImpl(
	brandRepo BrandRepo,
) *BrandSvcImpl {
	return &BrandSvcImpl{
		brandRepo: brandRepo,
	}
}

func (c *BrandSvcImpl) GetCriteria(filter *BrandFilter) ([]Criterion, error) {
	return c.brandRepo.GetCriteria(filter)
}
