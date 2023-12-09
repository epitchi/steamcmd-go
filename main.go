package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/epitchi/steamcmdgo/utils/fileserver"
	"github.com/epitchi/steamcmdgo/utils/steamcmd"
	"github.com/epitchi/steamcmdgo/utils/supabase"
	"github.com/epitchi/steamcmdgo/utils/zipfolder"
)

type SteamConfig struct {
	AppId    int    `json:"app_id"`
	AppName  string `json:"app_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// - SQL Query to Virtless, check lastest app_id -> if yes, execute, if not, wait for 12 hours
	gamesNeedDownload := supabase.SupabaseCheck("sontay.thinkmay.net", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ewogICJyb2xlIjogImFub24iLAogICJpc3MiOiAic3VwYWJhc2UiLAogICJpYXQiOiAxNjk0MDE5NjAwLAogICJleHAiOiAxODUxODcyNDAwCn0.EpUhNso-BMFvAJLjYbomIddyFfN--u-zCf0Swj9Ac6E")

	if reflect.DeepEqual(gamesNeedDownload, []supabase.Stores{}) {
		fmt.Println("No found any game need to download")
	}

	for _, p := range gamesNeedDownload {
		fmt.Println(p)
		foundSteamApp := findSteamInfo(p.Name)
		if foundSteamApp != nil {
			result := steamcmd.SteamCmdExec(
				foundSteamApp.AppId, 
				strings.Replace(foundSteamApp.AppName, " ", "_", -1), 
				foundSteamApp.Username, 
				foundSteamApp.Password,
			)
			
			if result != nil {
				zipfolder.ZipFile(result.SourceDir, "Games/" + result.SteamName)
			}

			port := flag.Int("port", 8080, "Port to listen on")
			directory := flag.String("dir", ".", "/Game")
			flag.Parse()
			err := fileserver.StartServer(*port, *directory)
			if err != nil {
				println("Error:", err)
				os.Exit(1)
			}

		}
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
		if p.AppName == targetAppName {
			return &p
		}
	}

	return nil //Not found
}
