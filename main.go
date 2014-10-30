package main

import (
	_ "github.com/Marktown/frontend/routers"
	"github.com/astaxie/beego"
)

func main() {
  beego.CopyRequestBody = true
	beego.Run()
}
