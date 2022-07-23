package repository

import (
	"context"
	"github.com/google/wire"
	"github.com/hackerchai/threatbook-ip-web/ent"
	"github.com/hackerchai/threatbook-ip-web/ent/threat"
)
import "github.com/hackerchai/threatbook-ip-web/internal/app/schema"

var ThreatRepoStruct = wire.Struct(new(ThreatRepo), "*")

type ThreatRepo struct {
	DB *ent.Client
}

func (t *ThreatRepo) QueryThreatsWithPagnition(ctx context.Context, pagnition schema.PaginationParam) ([]*ent.Threat, schema.PaginationResponse, error) {
	current := pagnition.GetCurrent()
	pageSize := pagnition.GetPageSize()
	offset := (current - 1) * pageSize

	threats, err := t.DB.Threat.Query().Order(ent.Asc(threat.FieldID)).Offset(offset).Limit(pageSize).All(ctx)
	if err != nil {
		return nil, schema.PaginationResponse{}, err
	}
	if total, err := t.DB.Threat.Query().Count(ctx); err != nil {
		return nil, schema.PaginationResponse{}, err
	} else {
		return threats, schema.PaginationResponse{
			Current:  current,
			Total:    total,
			PageSize: pageSize,
		}, nil
	}
}
