package main

import (
	"app/config"
	"app/routers"
	"fmt"
)

func main() {
	r := routers.SetRouters()
	err := r.Run(fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port))
	if err != nil {
		fmt.Print(err)
	}
}
