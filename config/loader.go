package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/appengine/log"

	configTypes "github.com/naurffxiv/clearingway/config/types"
)

type BotConfig struct {
	Encounters map[string]*configTypes.EncounterConfig `json:"encounters"`
	Menus      map[string]*configTypes.MenuConfig      `json:"menus"`
}

func (cfg *BotConfig) GetEncounters() map[string]*configTypes.EncounterConfig {
	return cfg.Encounters
}

func (cfg *BotConfig) GetMenus() map[string]*configTypes.MenuConfig {
	return cfg.Menus
}

func (cfg *BotConfig) GetEncounterByID(id int) *configTypes.EncounterConfig {
	for _, encounter := range cfg.Encounters {
		for _, encounterID := range encounter.IDs {
			if encounterID == id {
				return encounter
			}
		}
	}
	return nil
}

func (cfg *BotConfig) GetEncounterByName(name string) *configTypes.EncounterConfig {
	if encounter, ok := cfg.Encounters[name]; ok {
		return encounter
	}
	return nil
}

func (cfg *BotConfig) GetMenuByName(name string) *configTypes.MenuConfig {
	if menu, ok := cfg.Menus[name]; ok {
		return menu
	}
	return nil
}

func (cfg *BotConfig) parseMenuConfig(ctx context.Context, path string, data []byte) error {
	var menuConfig configTypes.MenuConfig
	if err := json.Unmarshal(data, &menuConfig); err != nil {
		log.Errorf(ctx, "error unmarshaling menu config file %s: %v", path, err)
		return err
	}
	if cfg.Menus == nil {
		cfg.Menus = make(map[string]*configTypes.MenuConfig)
	}
	cfg.Menus[menuConfig.Name] = &menuConfig
	return nil
}

func (cfg *BotConfig) parseEncounterConfig(ctx context.Context, path string, data []byte) error {
	var encounterConfig configTypes.EncounterConfig
	if err := json.Unmarshal(data, &encounterConfig); err != nil {
		log.Errorf(ctx, "error unmarshaling encounter config file %s: %v", path, err)
		return err
	}
	if cfg.Encounters == nil {
		cfg.Encounters = make(map[string]*configTypes.EncounterConfig)
	}
	cfg.Encounters[encounterConfig.Name] = &encounterConfig
	return nil
}

func (cfg *BotConfig) parseConfigFile(path string, info os.FileInfo, err error) error {
	ctx := context.Background()

	if err != nil {
		return err
	}

	// Skip non-JSON files and directories
	if info.IsDir() || filepath.Ext(path) != ".json" {
		return nil
	}

	file, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("error reading config file %s: %s", path, err)
	}

	lastDirInPath := filepath.Base(filepath.Dir(path))
	switch lastDirInPath {
	case "menus":
		if err := cfg.parseMenuConfig(ctx, path, file); err != nil {
			return err
		}
	case "ultimates":
	case "savages":
	case "extremes":
		if err := cfg.parseEncounterConfig(ctx, path, file); err != nil {
			return err
		}
	default:
		log.Errorf(ctx, "unknown config file type for file %s", path)
		return fmt.Errorf("unknown config file type for file %s", path)
	}

	return nil
}

func (cfg *BotConfig) Init(ctx context.Context, configDir string) {
	err := filepath.Walk(configDir, cfg.parseConfigFile)
	if err != nil {
		log.Errorf(ctx, "error walking config dir: %s", err)
	}
}

func main() {
	ctx := context.Background()
	cfg := &BotConfig{}
	cfg.Init(ctx, "./config/data")

	fmt.Printf("Loaded BotConfig: %+v\n", cfg)
}
