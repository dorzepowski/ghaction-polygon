package main

import (
	"4chain-ag/ghaction-polygon/api"

	"github.com/gin-gonic/gin"
)

//go:generate swag init -g cmd/main.go -d ../. -o ../docs
//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		petstore.swagger.io
//	@BasePath	/v2

func main() {
	r := gin.New()
	r.GET("/testapi/get-string-by-int/:some_id", api.GetStringByInt)
	r.GET("//testapi/get-struct-array-by-string/:some_id", api.GetStructArrayByString)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
