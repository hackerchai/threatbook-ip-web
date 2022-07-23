package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hackerchai/threatbook-ip-web/internal/app/schema"
	"github.com/hackerchai/threatbook-ip-web/pkg/errors"
	"github.com/hackerchai/threatbook-ip-web/pkg/response"
	"log"
	"net/http"
)

// Response success with status ok
func ResOK(c *fiber.Ctx) {
	ResSuccess(c, schema.StatusResult{Status: schema.OKStatus})
}

// Response data with list object
func ResList(c *fiber.Ctx, v interface{}) {
	ResSuccess(c, schema.ListResult{List: v})
}

// Response pagination data object
func ResPage(c *fiber.Ctx, v interface{}, pr *schema.PaginationResponse) error {
	list := schema.ListResult{
		List:       v,
		Pagination: pr,
	}
	return ResSuccess(c, list)
}

// Response data object
func ResSuccess(c *fiber.Ctx, v interface{}) error {
	return ResJSON(c, http.StatusOK, v)
}

// Response json data with status code
func ResJSON(c *fiber.Ctx, status int, v interface{}) error {
	return c.Status(status).JSON(response.Response{
		Code:       0,
		StatusCode: status,
		Message:    response.MessageSuccess,
		Data:       v,
	})
}

// Response error object and parse error status code
func ResError(c *fiber.Ctx, err error, status ...int) error {
	var res *response.ResponseError

	if err != nil {
		if e, ok := err.(*response.ResponseError); ok {
			res = e
		} else {
			res = response.UnWrapResponse(errors.ErrInternalServer)
			res.ERR = err
		}
	} else {
		res = response.UnWrapResponse(errors.ErrInternalServer)
	}

	if len(status) > 0 {
		res.StatusCode = status[0]
	}

	if err := res.ERR; err != nil {
		if res.Message == "" {
			res.Message = err.Error()
			return err
		}

		if status := res.StatusCode; status >= 400 && status < 500 {
			log.Println(err.Error())
		} else if status >= 500 {
			log.Println(err.Error())
		}
	}

	eitem := schema.ErrorItem{
		Code:    res.Code,
		Message: res.Message,
	}
	return ResJSON(c, res.StatusCode, schema.ErrorResult{Error: eitem})
}
