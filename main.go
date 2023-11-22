package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type SteamConfig struct {
	AppId 		int		`json:"app_id"`
	AppName		string 	`json:"app_name"`
	Username	string	`json:"username"`	
	Password	string	`json:"password"`
}

func main(){

	gamePath := "c:\\GameServer\\"

	steamAppID := 730

	fmt.Println("Callback Virtless change")
	// /** Callback and login
		// - SQL Query to Virtless, check lastest app_id -> if yes, execute, if not, wait for 12 hours

	
	foundSteamApp := findSteamInfo(steamAppID)

	if foundSteamApp != nil {
		fmt.Println(foundSteamApp)
		steamDownload := exec.Command("steamcmd", 
			"+force_install_dir", 
			gamePath + "json[app_id].app_name", 
			"+login", "username",  "password", 
			"+app_update", strconv.Itoa(foundSteamApp.AppId), 
			"+quit")
		
		fmt.Println(steamDownload)
	}



	fmt.Println("Zip and Ship")

}


func findSteamInfo(targetAppId int) *SteamConfig {

	fileContent, err := os.ReadFile("steam_app.json")
	if err != nil {
		fmt.Println("Error reading JSON file", err)
		return nil
	}
	var config []SteamConfig
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	for _, p := range config {
		if p.AppId == targetAppId {
			return &p
		}
	}

	return nil //Not found
}