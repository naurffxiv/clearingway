package main

import (
	"github.com/naurffxiv/clearingway/config"
	"github.com/naurffxiv/clearingway/internal/clearingway"
)

func Config() *clearingway.Clearingway {
	// ------------------- LOAD ENV -------------------
	_ = config.LoadEnv() // currently unused, will be used to setup clearingway

	// -------------- SETUP CLEARINGWAY ---------------
	// TODO: Setup clearingway from config/

	// ---------------- LOAD CONFIG -------------------
	// TODO: Load config ftom JSON files from config/

	// -------------- RETURN CLEARINGWAY --------------
	return nil // TODO: Return clearingway
}
