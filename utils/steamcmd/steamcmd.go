package steamcmd

import (
	"os"

	"github.com/jensvandewiel/gosteamcmd"
	"github.com/jensvandewiel/gosteamcmd/console"
)

type SteamCMDInterface struct{
	SteamId 	int
	SteamName	string
	SourceDir	string
}

func SteamCmdExec(steam_id int, steam_name string, username string, pwd string) *SteamCMDInterface {

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
	return &SteamCMDInterface{
		SteamId: steam_id,
		SteamName: steam_name,
		SourceDir: "/home/epitchi/steamcmd/" + steam_name,
	}
}
