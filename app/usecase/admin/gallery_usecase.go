package usecase

import (
	"context"
	"encoding/json"
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/datatables"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
)

type galleryUsecase struct {
	galleryRepo repository.GalleryRepository
	redis       cache.Rueidis
}

type GalleryUsecase interface {
	FindById(ctx context.Context, id int64) (sqlc.GetGalleryRow, error)
	Store(ctx context.Context, request request.CreateGalleryRequest) (sqlc.Gallery, error)
	Update(ctx context.Context, request request.UpdateGalleryRequest) error
	Destroy(ctx context.Context, id int64) error
	Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error)
	Publish(ctx context.Context) error
}

func NewGalleryUsecase(galleryRepo repository.GalleryRepository, redis cache.Rueidis) GalleryUsecase {
	return &galleryUsecase{
		galleryRepo: galleryRepo,
		redis:       redis,
	}
}

func (u *galleryUsecase) FindById(ctx context.Context, id int64) (sqlc.GetGalleryRow, error) {
	return u.galleryRepo.FindById(ctx, id)
}

func (u *galleryUsecase) Store(ctx context.Context, request request.CreateGalleryRequest) (sqlc.Gallery, error) {
	var gallery sqlc.CreateGalleryParams

	if err := smapping.FillStruct(&gallery, smapping.MapFields(&request)); err != nil {
		return sqlc.Gallery{}, err
	}

	result, err := u.galleryRepo.Insert(ctx, gallery)
	if err != nil {
		return sqlc.Gallery{}, err
	}

	return result, nil
}

func (u *galleryUsecase) Update(ctx context.Context, request request.UpdateGalleryRequest) error {
	var gallery sqlc.UpdateGalleryParams

	if err := smapping.FillStruct(&gallery, smapping.MapFields(&request)); err != nil {
		return err
	}

	gallery.Image = helper.StringToPointer(*gallery.Image)

	return u.galleryRepo.Update(ctx, gallery)
}

func (u *galleryUsecase) Destroy(ctx context.Context, id int64) error {
	_, err := u.galleryRepo.FindById(ctx, id)
	if err != nil {
		return err
	}

	return u.galleryRepo.Delete(ctx, id)
}

func (u *galleryUsecase) Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error) {
	var total, filtered int64
	var search = params.Search

	orderColumn := map[string]string{
		"1": "image",
		"2": "image_caption",
	}

	arg := sqlc.DatatablesGalleryParams{
		Column1: search,
		Column2: orderColumn[params.OrderColumn],
		Column3: params.OrderDirection,
		Limit:   int32(params.Length),
		Offset:  int32(params.Start),
	}

	galleries, err := u.galleryRepo.Datatables(ctx, arg)
	if err != nil {
		return nil, err
	}

	if galleries == nil {
		galleries = []sqlc.DatatablesGalleryRow{}
	}

	total, err = u.galleryRepo.Count(ctx, "")
	if err != nil {
		return nil, err
	}

	if search != "" {
		filtered, err = u.galleryRepo.Count(ctx, search)
		if err != nil {
			return nil, err
		}
	} else {
		filtered = total
	}

	return datatables.NewDataTableResponse(params.Draw, total, filtered, galleries), nil
}

func (u *galleryUsecase) Publish(ctx context.Context) error {
	galleries, err := u.galleryRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	galleryBytes, err := json.Marshal(galleries)
	if err != nil {
		return err
	}

	return u.redis.Set(constants.KeyGalleries, galleryBytes, 0)
}
