package network

import (
	"fmt"
	"memo-point-com/internal/database"
	"time"
)

func handleDatabaseConnection() (e error) {
	var nbTry int = 0
	var err error
	for nbTry < 3 {
		fmt.Println("Tentative {", (nbTry + 1), "} de connexion à la base de donnée...")
		database.DbInstance, err = database.ConnectToMariaDB()
		if err != nil {
			fmt.Println("Tentative de connexion à la base de donnée échouée, nouvelle tentative dans 5 secondes...")
			nbTry++
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	if nbTry == 3 {
		fmt.Println("Erreur de connexion à la base de données après 3 tentatives, arrêt du serveur...")
		return err
	}

	return nil
}

func InitServer() {
	err := handleDatabaseConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	sqlDB, err := database.DbInstance.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sqlDB.Close()

	InitRoutes()
}
