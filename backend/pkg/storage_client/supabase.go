package storageclient

import (
	"fmt"

	supabase_str "github.com/supabase-community/storage-go"
)

func NewSupaBaseInstance(projectRefID, key string) *supabase_str.Client {
	storageClient := supabase_str.NewClient(fmt.Sprintf("https://%s.supabase.co/storage/v1", projectRefID), key, nil)
	return storageClient
}
