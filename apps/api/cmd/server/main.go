package main


import (
    "github.com/wantedbear007/project_skim/internal/env"
    "github.com/wantedbear007/project_skim/internal/logging"
    _"github.com/wantedbear007/project_skim/internal/cache"
    "github.com/wantedbear007/project_skim/migrations"
)


func main() {

    env.LoadEnv()

    logging.InitLogger()
    // cache.InitCache()

    // zap.L().Info("hello world")

    // _ := env.DATABASE_URL.GetValue()

    migrations.RunMigrations()
    

}
