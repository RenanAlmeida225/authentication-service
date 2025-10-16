package main

import (
	"authentication-service/api/di"
	"authentication-service/api/server"
	"authentication-service/utils"
	"log"
)

func main() {
	utils.InitConfig()

	d := di.New()
	d.InjectDependencies("postgres")

	defer d.CloseDB()

	if err := d.Dig.Invoke(func(s *server.Server) { s.Listen() }); err != nil {
		log.Fatalf("erro ao subir server: %v", err)
	}
}
