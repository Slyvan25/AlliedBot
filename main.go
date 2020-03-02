package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load enviorment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// set enviorment variables
	version := os.Getenv("version")
	botID := os.Getenv("client_id")
	// log a fancy ascii art with botId and Version number
	logoArt := fmt.Sprintf(`
    =============================================================================
        _    _ _ _          _ ____        _   
       / \  | | (_) ___  __| | __ )  ___ | |_ 
      / _ \ | | | |/ _ \/ _  |  _ \ / _ \| __|
     / ___ \| | | |  __/ (_| | |_) | (_) | |_ 
    /_/   \_\_|_|_|\___|\__,_|____/ \___/ \__|
	
    Version: %s
    BotId: %s
    =============================================================================
   `, version, botID)
	print(logoArt)
}
