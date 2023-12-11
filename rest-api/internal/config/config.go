package config

import (
	"fmt"
	"time"
)

func InitTimezone() {
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		fmt.Println("Erreur lors du chargement du timezone:", err)
		return
	}

	time.Local = loc
}
