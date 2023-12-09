package supabase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Stores struct {
	Id 			int 	`json:"id"`
	Name		string	`json:"name"`
}


func SupabaseCheck(proj string, anon_key string) []Stores {

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