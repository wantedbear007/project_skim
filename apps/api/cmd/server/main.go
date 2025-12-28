package main


import (
    "go.uber.org/zap"
    "github.com/wantedbear007/project_skim/internal/env"
    "github.com/wantedbear007/project_skim/internal/logging"
    "github.com/wantedbear007/project_skim/internal/cache"
)


func main() {

    env.LoadEnv()

    logging.InitLogger()
    cache.InitCache()

    zap.L().Info("hello world")

    x := env.DATABASE_URL.GetValue()

    

    print("value is %x ", x)
}
