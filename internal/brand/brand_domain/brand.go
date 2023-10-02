package brand_domain

import (
	"github.com/michael-bc/list-full-record-backend/internal/common/common_common"
	"time"
)

type Brand struct {
	ID            int64
	Name          string
	Image         *common_common.Image
	Location      Location
	ProductsCount int
	AverageRating float32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
