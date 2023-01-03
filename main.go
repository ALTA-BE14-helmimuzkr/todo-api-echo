package main

import (
	"github.com/ALTA-BE14-helmimuzkr/config"
	"github.com/ALTA-BE14-helmimuzkr/database"
	"github.com/ALTA-BE14-helmimuzkr/model"
)

func main() {
	// Setup Config
	c := config.InitConfig()

	// Setup Dataabse
	db := database.OpenConnectionMysql(c)

	// Setup Migration
	model.Migrate(db)

}
