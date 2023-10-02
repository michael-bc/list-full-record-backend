package brand_http

import (
	"github.com/labstack/echo/v4"
	"github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"
	"github.com/michael-bc/list-full-record-backend/internal/common/common_domain"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/base_error"
	"github.com/samber/lo"
	"net/http"
)

type CriterionController struct {
	brandService brand_domain.BrandService
}

func NewCriterionController(
	criterionService brand_domain.BrandService,
) *CriterionController {
	return &CriterionController{
		brandService: criterionService,
	}
}

func (c *CriterionController) SetupRoutes(router *echo.Echo) {
	criterionGroup := router.Group("/criteria")
	criterionGroup.GET("", c.GetCriteria)
}

func (c *CriterionController) GetCriteria(ctx echo.Context) error {
	var req GetCriteriaRequest

	if err := ctx.Bind(&req); err != nil {
		return base_error.NewBadRequestError(common_domain.BrandDomain, err, nil)
	}

	filter := BrandFilterFromGetCriteriaRequest(req)

	criteria, err := c.brandService.GetCriteria(&filter)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, lo.Map(criteria, func(criterion brand_domain.Criterion, _ int) CriterionDTO {
		return CriteriaToDTO(criterion)
	}))
}
