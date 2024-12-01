package repository

import (
	"context"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
)

type galleryRepository struct {
	db *sqlc.Queries
}

type GalleryRepository interface {
	GetAll(ctx context.Context) ([]sqlc.GetAllGalleryRow, error)
	FindById(ctx context.Context, id int64) (sqlc.GetGalleryRow, error)
	Insert(ctx context.Context, request sqlc.CreateGalleryParams) (sqlc.Gallery, error)
	Update(ctx context.Context, request sqlc.UpdateGalleryParams) error
	Delete(ctx context.Context, id int64) error
	Datatables(ctx context.Context, arg sqlc.DatatablesGalleryParams) ([]sqlc.DatatablesGalleryRow, error)
	Count(ctx context.Context, search string) (int64, error)
}

func NewGalleryRepository(db *sqlc.Queries) GalleryRepository {
	return &galleryRepository{
		db: db,
	}
}

func (r *galleryRepository) GetAll(ctx context.Context) ([]sqlc.GetAllGalleryRow, error) {
	galleries, err := r.db.GetAllGallery(ctx)
	if err != nil {
		return nil, err
	}

	return galleries, nil
}

func (r *galleryRepository) FindById(ctx context.Context, id int64) (sqlc.GetGalleryRow, error) {
	gallery, err := r.db.GetGallery(ctx, id)
	if err != nil {
		return sqlc.GetGalleryRow{}, err
	}

	return gallery, nil
}

func (r *galleryRepository) Insert(ctx context.Context, request sqlc.CreateGalleryParams) (sqlc.Gallery, error) {
	gallery, err := r.db.CreateGallery(ctx, request)
	if err != nil {
		return sqlc.Gallery{}, err
	}

	return gallery, nil
}

func (r *galleryRepository) Update(ctx context.Context, request sqlc.UpdateGalleryParams) error {
	return r.db.UpdateGallery(ctx, request)
}

func (r *galleryRepository) Delete(ctx context.Context, id int64) error {
	return r.db.DeleteGallery(ctx, id)
}

func (r *galleryRepository) Datatables(ctx context.Context, arg sqlc.DatatablesGalleryParams) ([]sqlc.DatatablesGalleryRow, error) {
	galleries, err := r.db.DatatablesGallery(ctx, arg)
	if err != nil {
		return nil, err
	}

	return galleries, nil
}

func (r *galleryRepository) Count(ctx context.Context, search string) (int64, error) {
	count, err := r.db.CountGallery(ctx, search)
	if err != nil {
		return 0, err
	}

	return count, nil
}
