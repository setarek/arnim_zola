package main

import (
	"github.com/setarek/arnim_zola/cmd/server"
	"github.com/setarek/arnim_zola/config"
)

func main()  {
	config.InitConfig()
	server.Init()
}



