package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type BotConfig struct {
	Encounters map[string]*EncounterConfig `json:"encounters"`
	Menus      map[string]*MenuConfig      `json:"menus"`
}

func (cfg *BotConfig) GetEncounters() map[string]*EncounterConfig {
	return cfg.Encounters
}

func (cfg *BotConfig) GetMenus() map[string]*MenuConfig {
	return cfg.Menus
}

func (cfg *BotConfig) GetEncounterByID(id int) *EncounterConfig {
	for _, encounter := range cfg.Encounters {
		for _, encounterID := range encounter.IDs {
			if encounterID == id {
				return encounter
			}
		}
	}
	return nil
}

func (cfg *BotConfig) GetEncounterByName(name string) *EncounterConfig {
	if encounter, ok := cfg.Encounters[name]; ok {
		return encounter
	}
	return nil
}

func (cfg *BotConfig) GetMenuByName(name string) *MenuConfig {
	if menu, ok := cfg.Menus[name]; ok {
		return menu
	}
	return nil
}

// TODO: Adjust to work with future menu state
//func (cfg *BotConfig) parseMenuConfig(path string, data []byte) error {
//	var menuConfig MenuConfig
//	if err := json.Unmarshal(data, &menuConfig); err != nil {
//		return fmt.Errorf("error unmarshaling menu config file %s: %w", path, err)
//	}
//	if cfg.Menus == nil {
//		cfg.Menus = make(map[string]*MenuConfig)
//	}
//	cfg.Menus[menuConfig.Name] = &menuConfig
//	return nil
//}

func (cfg *BotConfig) parseEncounterConfig(path string, data []byte) error {
	var encounterConfig EncounterConfig
	if err := json.Unmarshal(data, &encounterConfig); err != nil {
		return fmt.Errorf("error unmarshaling encounter config file %s: %w", path, err)
	}
	if cfg.Encounters == nil {
		cfg.Encounters = make(map[string]*EncounterConfig)
	}
	cfg.Encounters[encounterConfig.Name] = &encounterConfig
	return nil
}

func (cfg *BotConfig) parseConfigFile(path string, info os.FileInfo, err error) error {
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
	// TODO: Adjust to work with future menu state
	//case "menus":
	//	if err := cfg.parseMenuConfig(path, file); err != nil {
	//		return err
	//	}
	case "ultimates":
	case "savages":
	case "extremes":
		if err := cfg.parseEncounterConfig(path, file); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown config file type for file %s", path)
	}

	return nil
}

func InitBotConfig(configDir string) (*BotConfig, error) {
	cfg := &BotConfig{}
	err := filepath.Walk(configDir, cfg.parseConfigFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
