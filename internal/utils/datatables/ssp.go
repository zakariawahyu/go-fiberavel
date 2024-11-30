package datatables

import (
	"github.com/gofiber/fiber/v2"
)

type DataTableParams struct {
	Draw           int    `json:"draw"`
	Search         string `json:"search[value]"`
	Start          int    `json:"start"`
	Length         int    `json:"length"`
	OrderColumn    string `json:"order[0][column]"`
	OrderDirection string `json:"order[0][dir]"`
}

type DataTableResponse struct {
	Draw            int         `json:"draw"`
	RecordsTotal    int64       `json:"recordsTotal"`
	RecordsFiltered int64       `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
}

func ParseDataTableParams(c *fiber.Ctx) (*DataTableParams, error) {
	params := &DataTableParams{}
	if err := c.QueryParser(params); err != nil {
		return nil, err
	}

	searchParams := c.Query("search[value]")
	orderColumnParams := c.Query("order[0][column]")
	OrderDirectionParams := c.Query("order[0][dir]")
	params.Search = searchParams
	params.OrderColumn = orderColumnParams
	params.OrderDirection = OrderDirectionParams

	return params, nil
}

func NewDataTableResponse(draw int, total int64, filtered int64, data interface{}) *DataTableResponse {
	return &DataTableResponse{
		Draw:            draw,
		RecordsTotal:    total,
		RecordsFiltered: filtered,
		Data:            data,
	}
}
