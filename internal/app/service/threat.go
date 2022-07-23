package service

import (
	"context"
	"github.com/google/wire"
	"github.com/hackerchai/threatbook-ip-web/ent"
	"github.com/hackerchai/threatbook-ip-web/internal/app/repository"
	"github.com/hackerchai/threatbook-ip-web/internal/app/schema"
)

var ThreatSrvStruct = wire.Struct(new(ThreatSrv), "*")

type ThreatSrv struct {
	ThreatRepo repository.ThreatRepo
}

func (s *ThreatSrv) QueryThreatsWithPagnition(ctx context.Context, pagnition schema.PaginationParam) ([]*ent.Threat, schema.PaginationResponse, error) {
	return s.ThreatRepo.QueryThreatsWithPagnition(ctx, pagnition)
}
