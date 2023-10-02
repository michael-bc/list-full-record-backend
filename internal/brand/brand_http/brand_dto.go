package brand_http

import (
	"github.com/michael-bc/list-full-record-backend/internal/common/common_common"
	"time"
)

type BrandDTO struct {
	ID            int64                `json:"id"`
	Name          string               `json:"name"`
	Image         *common_common.Image `json:"image"`
	Location      LocationDTO          `json:"location"`
	ProductsCount int                  `json:"productsCount"`
	AverageRating float32              `json:"averageRating"`
	CreatedAt     time.Time            `json:"createdAt"`
	UpdatedAt     time.Time            `json:"updatedAt"`
}
