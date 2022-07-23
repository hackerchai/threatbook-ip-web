//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/hackerchai/threatbook-ip-web/internal/app/handler"
	"github.com/hackerchai/threatbook-ip-web/internal/app/repository"
	"github.com/hackerchai/threatbook-ip-web/internal/app/service"
)

func InitializeEvent() *Injector {
	panic(wire.Build(
		repository.RepositorySet,
		service.ServiceSet,
		handler.HandlerSet,
		InitDatabase,
		New,
		InjectorSet,
	))
	return new(Injector)
}
