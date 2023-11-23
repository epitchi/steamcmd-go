package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/epitchi/steamcmdgo/utils"
)

type SteamConfig struct {
	AppId    int    `json:"app_id"`
	AppName  string `json:"app_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}


func main() {
	fmt.Println("Callback Virtless change")
	// - SQL Query to Virtless, check lastest app_id -> if yes, execute, if not, wait for 12 hours
	gamesNeedDownload := utils.SupabaseCheck("virtless.thinkmay.net", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ewogICJyb2xlIjogImFub24iLAogICJpc3MiOiAic3VwYWJhc2UiLAogICJpYXQiOiAxNjk0MDE5NjAwLAogICJleHAiOiAxODUxODcyNDAwCn0.EpUhNso-BMFvAJLjYbomIddyFfN--u-zCf0Swj9Ac6E")

	for _, p := range gamesNeedDownload {
		fmt.Println(p)

		foundSteamApp := findSteamInfo(p.Name)
		if foundSteamApp != nil {
			fmt.Println(foundSteamApp)
		}
	}

	result := utils.SteamCmdExec(268910, "Cuphead", "", "")


	if result != nil {
		utils.ZipFile(result.SourceDir, result.SteamName)
	}

}

func findSteamInfo(targetAppName string) *SteamConfig {

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
		if strings.Replace(p.AppName, " ", "_", -1) == targetAppName {
			return &p
		}
	}

	return nil //Not found
}
