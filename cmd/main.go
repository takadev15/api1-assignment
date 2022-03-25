package main

import (
	"log"

	"github.com/takadev15/go-restapi/internal/database"
	"github.com/takadev15/go-restapi/internal/routers"
)

func main() {
  database.DBinit()
  r := routers.OrderRoutes()
  log.Fatal(r.Run(":3000"))
}
