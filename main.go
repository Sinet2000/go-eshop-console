package main

import (
	"context"
	"fmt"
	"github.com/Sinet2000/go-eshop-console/config"
	"github.com/Sinet2000/go-eshop-console/handlers"
	"github.com/Sinet2000/go-eshop-console/internal/db"
	"github.com/Sinet2000/go-eshop-console/internal/utils"
	"log"
)

func main() {
	config.LoadConfig()
	//_, err = db.NewPgService()
	// if err != nil {
	//	log.Fatalf("Failed to connect to PostgreSQL Db: %v", err)
	// }

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoDbContext, err := db.NewMongoService(config.GetEnv("MONGO_DB_NAME"), ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer func() {
		if err := mongoDbContext.Close(); err != nil {
			log.Printf("Error during MongoDB closure: %v", err)
		}
	}()

	for {
		isAdmin, err := utils.Confirm("Are you admin? [y/n]: ")
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		if isAdmin {
			adminMenuHandler := handlers.NewAdminHandler(mongoDbContext.DB)
			adminMenuHandler.RunAdminMenu(ctx)
		} else {
			clientMenuHandler := handlers.NewClientHandler(mongoDbContext.DB)
			clientMenuHandler.RunClientMenu(ctx)
		}

		shouldExitProgram := false
		for {
			shouldExitProgram, err = utils.Confirm("Do you want to exit the program? [y/n]: ")
			if err != nil {
				log.Fatalf("Error: %v", err)
			}

			break
		}

		if shouldExitProgram {
			break
		}
	}
}
