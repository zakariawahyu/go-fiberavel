package usecase

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tidwall/gjson"
	"github.com/zakariawahyu/go-fiberavel/internal/infrastructure/cache"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/constants"
	"github.com/zakariawahyu/go-fiberavel/internal/utils/helper"
	"sync"
)

type homeUsecase struct {
	redis cache.Rueidis
}

type HomeUsecase interface {
	GetConfigs(ctx context.Context, wg *sync.WaitGroup, configChan chan map[string]interface{}, errChan chan error)
	GetData(wg *sync.WaitGroup, dataChan chan map[string]interface{}, errChan chan error)
	Hget(guest string, wg *sync.WaitGroup, dataChan chan map[string]interface{}, errChan chan error)
}

func NewHomeUsecase(redis cache.Rueidis) HomeUsecase {
	return &homeUsecase{
		redis: redis,
	}
}

func (u *homeUsecase) GetConfigs(ctx context.Context, wg *sync.WaitGroup, configChan chan map[string]interface{}, errChan chan error) {
	wg.Add(1)
	defer func() {
		close(configChan)
		close(errChan)
		wg.Done()
	}()

	log.Infof("Start get configs")

	configs, err := u.redis.HGetAll(ctx, "configs")
	if err != nil {
		configChan <- nil
		errChan <- err
		return
	}

	event := gjson.Parse(configs["event"]).Map()
	eventDetail := fiber.Map{
		"title":         event["title"].String(),
		"description":   event["description"].String(),
		"image":         event["image"].String(),
		"image_caption": event["image_caption"].String(),
		"custom_data": fiber.Map{
			"date": helper.ParseUTC(event["custom_data"].Get("date").String()),
		},
		"is_active": event["is_active"].Bool(),
	}

	data := fiber.Map{
		"meta":  gjson.Parse(configs["meta"]).Value(),
		"cover": gjson.Parse(configs["cover"]).Value(),
		"event": eventDetail,
		"story": gjson.Parse(configs["story"]).Value(),
		"venue": gjson.Parse(configs["venue"]).Value(),
		"rsvp":  gjson.Parse(configs["rsvp"]).Value(),
		"gift":  gjson.Parse(configs["gift"]).Value(),
		"wish":  gjson.Parse(configs["wish"]).Value(),
		"thank": gjson.Parse(configs["thank"]).Value(),
	}

	configChan <- data
	errChan <- nil
	return
}

func (u *homeUsecase) GetData(wg *sync.WaitGroup, dataChan chan map[string]interface{}, errChan chan error) {
	wg.Add(1)
	defer func() {
		close(dataChan)
		close(errChan)
		wg.Done()
	}()

	log.Infof("Start get data")

	resCouples, err := u.redis.Get(constants.KeyCouples)
	if err != nil {
		dataChan <- nil
		errChan <- err
		return
	}

	resVenueDetails, err := u.redis.Get(constants.KeyVenues)
	if err != nil {
		dataChan <- nil
		errChan <- err
		return
	}

	venues := gjson.Parse(string(resVenueDetails)).Array()
	VenueDetails := make(map[int]interface{})
	for key, value := range venues {
		jam, hari, tanggal := helper.ParseDate(value.Get("date_held").String())
		VenueDetails[key] = fiber.Map{
			"id":       value.Get("id").Int(),
			"name":     value.Get("name").String(),
			"location": value.Get("location").String(),
			"address":  value.Get("address").String(),
			"map":      value.Get("map").String(),
			"jam":      jam,
			"hari":     hari,
			"tanggal":  tanggal,
		}
	}

	resGalleries, err := u.redis.Get(constants.KeyGalleries)
	if err != nil {
		dataChan <- nil
		errChan <- err
		return
	}

	resGifts, err := u.redis.Get(constants.KeyGift)
	if err != nil {
		dataChan <- nil
		errChan <- err
		return
	}

	data := fiber.Map{
		"couples":       gjson.Parse(string(resCouples)).Value(),
		"venue_details": VenueDetails,
		"galleries":     gjson.Parse(string(resGalleries)).Value(),
		"gifts":         gjson.Parse(string(resGifts)).Value(),
	}

	dataChan <- data
	errChan <- nil
	return
}

func (u *homeUsecase) Hget(guest string, wg *sync.WaitGroup, dataChan chan map[string]interface{}, errChan chan error) {
	wg.Add(1)
	defer func() {
		close(dataChan)
		close(errChan)
		wg.Done()
	}()

	log.Infof("Start hget")

	res, err := u.redis.HGet(constants.KeyGuests, guest)
	if err != nil {
		dataChan <- nil
		errChan <- err
		return
	}

	data := fiber.Map{
		"guest": gjson.Parse(res).Value(),
	}

	dataChan <- data
	errChan <- nil
	return
}
