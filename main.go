package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/jensvandewiel/gosteamcmd"
	"github.com/jensvandewiel/gosteamcmd/console"

)

type SteamConfig struct {
	AppId 		int		`json:"app_id"`
	AppName		string 	`json:"app_name"`
	Username	string	`json:"username"`	
	Password	string	`json:"password"`
}

type Stores struct {
	Id 			int 	`json:"id"`
	Name		string	`json:"name"`
}

func main(){
	fmt.Println("Callback Virtless change")
	// - SQL Query to Virtless, check lastest app_id -> if yes, execute, if not, wait for 12 hours
	// gamesNeedDownload := supabaseCheck("virtless.thinkmay.net", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ewogICJyb2xlIjogImFub24iLAogICJpc3MiOiAic3VwYWJhc2UiLAogICJpYXQiOiAxNjk0MDE5NjAwLAogICJleHAiOiAxODUxODcyNDAwCn0.EpUhNso-BMFvAJLjYbomIddyFfN--u-zCf0Swj9Ac6E")



	// for _, p := range gamesNeedDownload {
	// 	fmt.Println(p)

	// 	foundSteamApp := findSteamInfo(p.Name)
	// 	if foundSteamApp != nil {
	// 		fmt.Println(foundSteamApp)
	// 		// steamDownload := exec.Command("steamcmd", 
	// 		// 	"+force_install_dir", 	gamePath + "json[app_id].app_name", 
	// 		// 	"+login", 				"username",  "password", 
	// 		// 	"+app_update", 			strconv.Itoa(foundSteamApp.AppId), 
	// 		// 	"+quit")
	// 	}
	// }



	steamcmd := SteamCmdExec(2430930, "rust_server", "", "")
	fmt.Println(steamcmd)


	fmt.Println("Zip and Ship")

}

func SteamCmdExec(steam_id int, steam_name string, username string, pwd string) *string {
	fmt.Println("hehe")

	prompts := []*gosteamcmd.Prompt{
		gosteamcmd.ForceInstallDir("/home/epitchi/steamcmd/" + steam_name),
		gosteamcmd.Login(username, pwd, ""),
		gosteamcmd.AppUpdate(steam_id, "", true),
	}


	cmd := gosteamcmd.New(os.Stdout, prompts, "/usr/games/steamcmd")

	cmd.Console.Parser.OnInformationReceived = func(action console.Action, progress float64, currentWritten, total uint64) {
		println("")
	}
	cmd.Console.Parser.OnAppInstalled = func(app uint32) {
		println("App installed: ", app, " Yay!")
	}

	_, err := cmd.Run()

	if err != nil {
		panic(err)
	}
	return nil
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

func supabaseCheck(proj string, anon_key string) []Stores {

	previous12Hour := time.Now().Add(-12 * time.Hour).UTC();

	req,err := http.NewRequest("GET",

	// TODO: rpc check pending games to download
		fmt.Sprintf("https://%s/rest/v1/stores?select=id,name&created_at=gt.%s", proj, previous12Hour.Format(time.RFC3339)),
		bytes.NewBuffer([]byte("")))
	if err != nil {
		panic(err)
	}

	fmt.Println( )

	req.Header.Set("apikey", anon_key)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s",anon_key))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	} else if resp.StatusCode != 200 {
		panic("unable to fetch constant from server")
	}

	body, _ := io.ReadAll(resp.Body)

	var data [](Stores)
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return data
} 