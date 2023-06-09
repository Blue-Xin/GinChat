package main

import (
	"GinChat/router"
	"GinChat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()

	r.Run(":8848") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
