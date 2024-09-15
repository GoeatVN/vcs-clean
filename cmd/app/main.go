package main

import (
	"github.com/resotto/goilerplate/internal/app/adapter"
	"github.com/spf13/viper"
)

func main() {
	r := adapter.Router()

	viper.SetDefault("BITBANK_HOST", "https://public.bitbank.cc")
	// for local development
	viper.SetDefault("PGHOST", "51.79.212.76")
	viper.SetDefault("PGUSER", "postgres")
	viper.SetDefault("PGPASSWORD", "kD9gdhQh")
	viper.SetDefault("PGPORT", "15432")
	viper.SetDefault("PGDATABASE", "vcs")
	viper.BindEnv("PGHOST", "PGHOST")
	viper.BindEnv("PGUSER", "PGUSER")
	viper.BindEnv("PGPASSWORD", "PGPASSWORD")
	viper.BindEnv("PGPORT", "PGPORT")
	viper.BindEnv("PGDATABASE", "PGDATABASE")
	r.Run(":8080")
}
