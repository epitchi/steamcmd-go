package supabase

import (
	"fmt"
	"testing"
)

func TestSupabase(t *testing.T){
	gamesNeedDownload := SupabaseCheck("sontay.thinkmay.net", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ewogICJyb2xlIjogImFub24iLAogICJpc3MiOiAic3VwYWJhc2UiLAogICJpYXQiOiAxNjk0MDE5NjAwLAogICJleHAiOiAxODUxODcyNDAwCn0.EpUhNso-BMFvAJLjYbomIddyFfN--u-zCf0Swj9Ac6E")

	fmt.Print(gamesNeedDownload)
}