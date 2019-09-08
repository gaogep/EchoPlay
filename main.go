package main

import (
	"github.com/gaogep/EchoPlay/routers"
)

func main() {
	e := routers.InitRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
