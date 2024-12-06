package usecase

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-fiberavel/app/http/request"
	repository "github.com/zakariawahyu/go-fiberavel/app/repository/admin"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	sqlc "github.com/zakariawahyu/go-fiberavel/internal/sqlc/generated"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/datatables"
)

type guestUsecase struct {
	guestRepo repository.GuestRepository
	redis     cache.Rueidis
}

type GuestUsecase interface {
	FindById(ctx context.Context, id int64) (sqlc.GetGuestRow, error)
	Store(ctx context.Context, request request.CreateGuestRequest) error
	Update(ctx context.Context, request request.UpdateGuestRequest) error
	Destroy(ctx context.Context, id int64) error
	Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error)
	Publish(ctx context.Context) error
}

func NewGuestUsecase(guestRepo repository.GuestRepository, redis cache.Rueidis) GuestUsecase {
	return &guestUsecase{
		guestRepo: guestRepo,
		redis:     redis,
	}
}

func (u *guestUsecase) FindById(ctx context.Context, id int64) (sqlc.GetGuestRow, error) {
	return u.guestRepo.FindById(ctx, id)
}

func (u *guestUsecase) Store(ctx context.Context, request request.CreateGuestRequest) error {
	var guest sqlc.CreateGuestParams

	if err := smapping.FillStruct(&guest, smapping.MapFields(&request)); err != nil {
		return err
	}

	return u.guestRepo.Insert(ctx, guest)
}

func (u *guestUsecase) Update(ctx context.Context, request request.UpdateGuestRequest) error {
	var guest sqlc.UpdateGuestParams

	if err := smapping.FillStruct(&guest, smapping.MapFields(&request)); err != nil {
		return err
	}

	return u.guestRepo.Update(ctx, guest)
}

func (u *guestUsecase) Destroy(ctx context.Context, id int64) error {
	_, err := u.guestRepo.FindById(ctx, id)
	if err != nil {
		return err
	}

	return u.guestRepo.Delete(ctx, id)
}

func (u *guestUsecase) Datatables(ctx context.Context, params *datatables.DataTableParams) (*datatables.DataTableResponse, error) {
	var total, filtered int64
	var search = params.Search

	orderColumn := map[string]string{
		"1": "name",
		"2": "slug",
		"3": "is_gift",
	}

	arg := sqlc.DatatablesGuestParams{
		Column1: search,
		Column2: orderColumn[params.OrderColumn],
		Column3: params.OrderDirection,
		Limit:   int32(params.Length),
		Offset:  int32(params.Start),
	}

	guests, err := u.guestRepo.Datatables(ctx, arg)
	if err != nil {
		return nil, err
	}

	if guests == nil {
		guests = []sqlc.DatatablesGuestRow{}
	}

	total, err = u.guestRepo.Count(ctx, "")
	if err != nil {
		return nil, err
	}

	if search != "" {
		filtered, err = u.guestRepo.Count(ctx, search)
		if err != nil {
			return nil, err
		}
	} else {
		filtered = total
	}

	return datatables.NewDataTableResponse(params.Draw, total, filtered, guests), nil
}

func (u *guestUsecase) Publish(ctx context.Context) error {
	guests, err := u.guestRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, guest := range guests {
		log.Infof("Guest: %v", guest)
		guestBytes, err := json.Marshal(guest)
		if err != nil {
			return err
		}

		if err := u.redis.Hset(constants.KeyGuests, guest.Slug, string(guestBytes)); err != nil {
			return err
		}
	}

	return nil
}
