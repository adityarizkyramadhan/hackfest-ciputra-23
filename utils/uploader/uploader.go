package uploader

import (
	"os"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

var SupClient = supabasestorageuploader.NewSupabaseClient(
	os.Getenv("SUPABASE_PROJECT_URL"),
	os.Getenv("SUPABASE_PROJECT_API_KEYS"),
	os.Getenv("SUPABASE_STORAGE_NAME"),
	os.Getenv("SUPABASE_STORAGE_FOLDER"),
)
