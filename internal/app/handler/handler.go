package handler

import "github.com/google/wire"

var HandlerStruct = wire.Struct(new(Handler), "*")

var HandlerSet = wire.NewSet(HandlerStruct, ThreatAPIStruct)

type Handler struct {
	ThreatAPI *ThreatAPI
}
