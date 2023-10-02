package brand_http

import "github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"

func CriteriaToDTO(criterion brand_domain.Criterion) CriterionDTO {
	return CriterionDTO{
		Category: CategoryToDTO(criterion.Category),
		Location: LocationToDTO(criterion.Location),
		Brands:   BrandsToDTOs(criterion.Brands),
	}
}
