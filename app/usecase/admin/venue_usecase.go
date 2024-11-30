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
)

type venueUsecase struct {
	venueRepo repository.VenueRepository
	redis     cache.Rueidis
}

type VenueUsecase interface {
	FindById(ctx context.Context, id int64) (sqlc.GetVenueRow, error)
	Store(ctx context.Context, request request.CreateVenueRequest) (sqlc.Venue, error)
	Update(ctx context.Context, request request.UpdateVenueRequest) error
	Destroy(ctx context.Context, id int64) error
	Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error)
	Publish(ctx context.Context) error
}

func NewVenueUsecase(venueRepo repository.VenueRepository, redis cache.Rueidis) VenueUsecase {
	return &venueUsecase{
		venueRepo: venueRepo,
		redis:     redis,
	}
}

func (u *venueUsecase) FindById(ctx context.Context, id int64) (sqlc.GetVenueRow, error) {
	return u.venueRepo.FindById(ctx, id)
}

func (u *venueUsecase) Store(ctx context.Context, request request.CreateVenueRequest) (sqlc.Venue, error) {
	var venue sqlc.CreateVenueParams

	if err := smapping.FillStruct(&venue, smapping.MapFields(&request)); err != nil {
		return sqlc.Venue{}, err
	}

	result, err := u.venueRepo.Insert(ctx, venue)
	if err != nil {
		return sqlc.Venue{}, err
	}

	return result, nil
}

func (u *venueUsecase) Update(ctx context.Context, request request.UpdateVenueRequest) error {
	var venue sqlc.UpdateVenueParams

	if err := smapping.FillStruct(&venue, smapping.MapFields(&request)); err != nil {
		return err
	}

	return u.venueRepo.Update(ctx, venue)
}

func (u *venueUsecase) Destroy(ctx context.Context, id int64) error {
	_, err := u.venueRepo.FindById(ctx, id)
	if err != nil {
		return err
	}

	return u.venueRepo.Delete(ctx, id)
}

func (u *venueUsecase) Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error) {
	var total, filtered int64
	var search = params.Search

	orderColumn := map[string]string{
		"1": "name",
		"2": "location",
		"3": "date_held",
	}

	arg := sqlc.DatatablesVenueParams{
		Column1: search,
		Column2: orderColumn[params.OrderColumn],
		Column3: params.OrderDirection,
		Limit:   int32(params.Length),
		Offset:  int32(params.Start),
	}

	venues, err := u.venueRepo.Datatables(ctx, arg)
	if err != nil {
		return nil, err
	}

	if venues == nil {
		venues = []sqlc.DatatablesVenueRow{}
	}

	total, err = u.venueRepo.Count(ctx, "")
	if err != nil {
		return nil, err
	}

	if search != "" {
		filtered, err = u.venueRepo.Count(ctx, search)
		if err != nil {
			return nil, err
		}
	} else {
		filtered = total
	}

	return datatables.NewDataTableResponse(params.Draw, total, filtered, venues), nil
}

func (u *venueUsecase) Publish(ctx context.Context) error {
	venues, err := u.venueRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	venueBytes, err := json.Marshal(venues)
	if err != nil {
		return err
	}

	return u.redis.Set(constants.KeyVenues, venueBytes, 0)
}
