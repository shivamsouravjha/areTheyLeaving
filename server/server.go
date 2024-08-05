package server

import (
	"areTheyLeaving/helper/mongo"
	"areTheyLeaving/routes"
)

func Init() {
	mongo.InitDb()
	r := routes.NewRouter()
	r.Run(":8080")
}
