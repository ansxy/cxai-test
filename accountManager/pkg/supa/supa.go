package supa

import (
	"github.com/ansxy/golang-boilerplate-gin/config"
	"github.com/supabase-community/supabase-go"
)

func NewSupaClient(cnf *config.SupabaseConfig) (*supabase.Client, error) {
	client, err := supabase.NewClient(cnf.ProjectRef, cnf.ApiKey, nil)
	if err != nil {
		return nil, err
	}

	return client, err
}
