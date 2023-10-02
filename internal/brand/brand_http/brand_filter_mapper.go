package brand_http

import "github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"

func BrandFilterFromGetCriteriaRequest(request GetCriteriaRequest) brand_domain.BrandFilter {
	return brand_domain.BrandFilter{
		Search: request.Search,
	}
}
