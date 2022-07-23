package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/hackerchai/threatbook-ip-web/internal/app/schema"
	"github.com/hackerchai/threatbook-ip-web/internal/app/service"
	"github.com/hackerchai/threatbook-ip-web/internal/app/web"
	"github.com/hackerchai/threatbook-ip-web/pkg/errors"
	"log"
)

var ThreatAPIStruct = wire.Struct(new(ThreatAPI), "*")

type ThreatAPI struct {
	ThreatSrv *service.ThreatSrv
}

// @Tags Threat
// @Summary 获取所有的威胁IP信息
// @Description 获取所有的威胁IP信息, 并分页的形式呈现
// @ID get-threats
// @Accept  json
// @Produce  json
// @Param current query integer false "当前查询页码(默认 1)"
// @Param page_sizet query integer false "当前查询每页显示条数(默认 10)"
// @Success 200 {object} schema.ThreatsResponse
// @Failure 400 {object} schema.ErrorResponse "{code:1002,status_code:401,message:no permission}"
// @Failure 500 {object} schema.ErrorResponse "{code:1001,status_code:504,message:internal server error}"
// @Router /api/threats [get]
func (h *ThreatAPI) GetThreats(c *fiber.Ctx) error {
	ctx := c.Context()
	var pagnitionParam schema.PaginationParam
	if err := c.QueryParser(&pagnitionParam); err != nil {
		log.Println("Parse query failed: ", err)
		return web.ResError(c, errors.ErrInvalidParameter)
	}
	threats, pagnitionParamResp, err := h.ThreatSrv.QueryThreatsWithPagnition(ctx, pagnitionParam)
	if err != nil {
		log.Println("Database query failed: ", err)
		err := web.ResError(c, err)
		if err != nil {
			return err
		}
	}
	return web.ResPage(c, threats, &pagnitionParamResp)
}
