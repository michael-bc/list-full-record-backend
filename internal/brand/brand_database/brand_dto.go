package brand_database

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/michael-bc/list-full-record-backend/internal/common/common_common"
	"time"
)

type brandDTO struct {
	ID            int64                `json:"id"`
	Name          string               `json:"name"`
	Image         *common_common.Image `json:"image"`
	Location      locationDTO          `json:"location"`
	ProductsCount int                  `json:"productsCount"`
	AverageRating float32              `json:"averageRating"`
	CreatedAt     time.Time            `json:"createdAt"`
	UpdatedAt     time.Time            `json:"updatedAt"`
}

type brandDTOs []brandDTO

func (b brandDTOs) Value() (driver.Value, error) {
	return json.Marshal(b)
}

func (b *brandDTOs) Scan(data any) error {
	if data == nil {
		return nil
	}

	switch val := data.(type) {
	case []byte:
		return json.Unmarshal(val, &b)
	case string:
		return json.Unmarshal([]byte(val), &b)
	default:
		return fmt.Errorf("unexpected data type - %T", data)
	}
}
