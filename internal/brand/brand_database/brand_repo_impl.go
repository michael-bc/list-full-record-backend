package brand_database

import (
	"database/sql"
	"fmt"
	"github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"
	"github.com/michael-bc/list-full-record-backend/internal/common/common_domain"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/base_error"
	"github.com/samber/lo"
)

type BrandRepoImpl struct {
	db *sql.DB
}

func NewBrandRepoImpl(db *sql.DB) *BrandRepoImpl {
	return &BrandRepoImpl{db: db}
}

func (r *BrandRepoImpl) scanCriteria(rows *sql.Rows) ([]criterionDTO, error) {
	var criteria []criterionDTO

	for rows.Next() {
		criterion := criterionDTO{}

		if err := rows.Scan(
			&criterion.Category,
			&criterion.Location,
			&criterion.Brands,
		); err != nil {
			return nil, base_error.NewDatabaseError(common_domain.BrandDomain, err)
		}

		criteria = append(criteria, criterion)
	}

	return criteria, nil
}

func (r *BrandRepoImpl) GetCriteria(filter *brand_domain.BrandFilter) ([]brand_domain.Criterion, error) {
	var where string
	var values []any

	if filter != nil && filter.Search != "" {
		where += fmt.Sprintf(
			`WHERE CONCAT(p.name, b.name, c.name, l.city) ILIKE $1`)
		values = append(values, fmt.Sprint("%", filter.Search, "%"))
	}

	rows, err := r.db.Query(fmt.Sprintf(`
		SELECT
			JSON_BUILD_OBJECT(
				'id', c.id,
				'name', c.name,
				'createdAt', CAST(c.created_at AS TIMESTAMP) AT time zone 'UTC',
				'updatedAt', CAST(c.updated_at AS TIMESTAMP) AT time zone 'UTC'
			) as category,
			JSON_BUILD_OBJECT(
				'id', l.id,
				'city', l.city,
				'createdAt', CAST(l.created_at AS TIMESTAMP) AT time zone 'UTC',
				'updatedAt', CAST(l.updated_at AS TIMESTAMP) AT time zone 'UTC'
			) as location,
			JSON_AGG(
				JSON_BUILD_OBJECT(
						'id', b.id,
						'name', b.name,
						'image', json_build_object(
							'url', b.image_url
						),
						'location', JSON_BUILD_OBJECT(
							'id', l.id,
							'city', l.city,
							'createdAt', CAST(l.created_at AS TIMESTAMP) AT time zone 'UTC',
							'updatedAt', CAST(l.updated_at AS TIMESTAMP) AT time zone 'UTC'
						),
						'productsCount', (SELECT COUNT(*) FROM products AS p WHERE p.brand_id = b.id),
						'averageRating', (SELECT AVG(p.rating) FROM products AS p WHERE p.brand_id = b.id),
						'createdAt', CAST(b.created_at AS TIMESTAMP) AT time zone 'UTC',
						'updatedAt', CAST(b.updated_at AS TIMESTAMP) AT time zone 'UTC'
				)
			) as brands
		FROM brands as b
			INNER JOIN locations l ON b.location_id = l.id
			INNER JOIN products p ON b.id = p.brand_id
			INNER JOIN categories c ON p.category_id = c.id
		%s
		GROUP BY c.id, l.id
	`, where), values...)
	if err != nil {
		return nil, base_error.NewDatabaseError(common_domain.BrandDomain, err)
	}

	dtos, err := r.scanCriteria(rows)
	if err != nil {
		return nil, err
	}

	return lo.Map(dtos, func(dto criterionDTO, _ int) brand_domain.Criterion {
		return criterionFromDTO(dto)
	}), nil
}
