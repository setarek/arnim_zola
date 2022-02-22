package server

import (
	"fmt"
	"github.com/setarek/arnim_zola/config"
	"github.com/setarek/arnim_zola/internal/router"
)

func Init() {
	config := config.GetConfig()
	r := router.NewRouter()
	r.Run(fmt.Sprintf("%s:%s", config.GetString("hostname"), config.GetString("port")))
}
