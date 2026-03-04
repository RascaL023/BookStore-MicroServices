
package repository

import (
	"context"
	"time"
	"writer/internal/cache"
	"writer/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type writerRepository struct { db *pgxpool.Pool }

func New(db *pgxpool.Pool) WriterRepository {
	return &writerRepository{db: db}
}



func (r *writerRepository) Save(
	ctx context.Context, w model.Writer,
) (*model.Writer, error) {
	ret := " RETURNING id, name, city, email;"
	var err error

	if w.Id == 0 {
		q := `
			INSERT INTO writers (name, city, email) 
			VALUES ($1, $2, $3)
		` + ret

		err = r.db.QueryRow(
			ctx, q,
			w.Name,
			w.City,
			w.Email,
		).Scan(&w.Id, &w.Name, &w.City, &w.Email)
	} else {
		q := `
			UPDATE writers 
			SET 
				name = $2, city = $3, 
				email = $4, is_active = $5
			WHERE id = $1 
		` + ret

		err = r.db.QueryRow(
			ctx, q,
			w.Id,
			w.Name,
			w.City,
			w.Email,
			w.IsActive,
		).Scan(&w.Id, &w.Name, &w.City, &w.Email)
	}

	if err != nil { return nil, err }
	return &w, nil
}


// ==================== GET DATA ====================
func (r *writerRepository) FindByID(
	ctx context.Context, id int64,
) (*model.Writer, error) {
	w := &model.Writer{}
	q := `
		SELECT id, name, city, email 
		FROM writers WHERE id = $1
	`
	err := r.db.QueryRow(ctx, q, id).
		Scan(&w.Id, &w.Name, &w.City, &w.Email)

	if err != nil { return nil, err }
	return w, nil
}

func (r *writerRepository) FindAll(
	ctx context.Context,
	offset, size int64,
) ([]*model.Writer, error) {
	q := `
		SELECT id, name, city, email, is_active 
		FROM writers LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, q, size, offset)
	if err != nil { return nil, err }
	defer rows.Close()
	
	var list []*model.Writer
	for rows.Next() {
		var temp model.Writer
		err := rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.City,
			&temp.Email,
			&temp.IsActive,
		)

		if err != nil { return nil, err }
		list = append(list, &temp)
	}

	if err := rows.Err(); err != nil { return  nil, err }
	return list, nil
}

func (r *writerRepository) Count(ctx context.Context) (
	total int64, err error,
) {
	q := `SELECT COUNT(*) FROM writers`
	err = r.db.QueryRow(ctx, q).Scan(&total)
	if err != nil { return 0, err }

	cache.Client.Set(ctx, "writer:count", total, 10*time.Minute)
	return total, nil
}

func (r *writerRepository) FindByName(
	ctx context.Context, 
	name string,
	offset, size int64,
) ([]*model.Writer, error) {
	q := `
		SELECT id, name, city, email, is_active 
		FROM writers
		WHERE name ILIKE $1 LIMIT $2 OFFSET $3
	`
	
	name = "%" + name + "%"
	rows, err := r.db.Query(ctx, q, name, size, offset)
	if err != nil { return nil, err }
	defer rows.Close()
	
	var list []*model.Writer
	for rows.Next() {
		var temp model.Writer
		err := rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.City,
			&temp.Email,
			&temp.IsActive,
		)

		if err != nil { return nil, err }
		list = append(list, &temp)
	}

	if err := rows.Err(); err != nil { return  nil, err }
	return list, nil
}
// ==================== GET DATA ====================
