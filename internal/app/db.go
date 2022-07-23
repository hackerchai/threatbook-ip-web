package app

import (
	"github.com/hackerchai/threatbook-ip-web/ent"
	"github.com/hackerchai/threatbook-ip-web/internal/app/global"
	_ "github.com/lib/pq"
	"log"
)

func InitDatabase() *ent.Client {
	deployMode := global.CONFIG.Common.DeployMode

	var dbCfg = ""
	switch deployMode {
	case "development":
		dbCfg = global.CONFIG.DevDatabase.Dsn()
	case "production":
		dbCfg = global.CONFIG.ProdDatabase.Dsn()
	default:
		log.Fatal("unsupported deploy mode: " + deployMode)
		return nil
	}
	client, err := ent.Open("postgres", dbCfg)
	if err != nil {
		log.Fatal("database connection error: " + err.Error())
		return nil
	}
	return client
}
