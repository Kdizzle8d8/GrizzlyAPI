package database

import (
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
	"os"
)

func GetClient() *supa.Client {
	godotenv.Load("../.env")
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseAnonKey := os.Getenv("SUPABASE_ANON_KEY")

	sb := supa.CreateClient(supabaseURL, supabaseAnonKey)
	return sb
}
