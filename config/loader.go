package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	configTypes "github.com/naurffxiv/clearingway/config/types"
)

type BotConfig struct {
	Encounters map[string]*configTypes.EncounterConfig `json:"encounters"`
	Menus      map[string]*configTypes.MenuConfig      `json:"menus"`
}

func (bc *BotConfig) GetEncounters() map[string]*configTypes.EncounterConfig {
	return bc.Encounters
}

func (bc *BotConfig) GetMenus() map[string]*configTypes.MenuConfig {
	return bc.Menus
}

func (bc *BotConfig) GetEncounterByID(id int) *configTypes.EncounterConfig {
	for _, encounter := range bc.Encounters {
		for _, encounterID := range encounter.IDs {
			if encounterID == id {
				return encounter
			}
		}
	}
	return nil
}

func (bc *BotConfig) GetEncounterByName(name string) *configTypes.EncounterConfig {
	if encounter, ok := bc.Encounters[name]; ok {
		return encounter
	}
	return nil
}

func (bc *BotConfig) GetMenuByName(name string) *configTypes.MenuConfig {
	if menu, ok := bc.Menus[name]; ok {
		return menu
	}
	return nil
}

func (bc *BotConfig) parseMenuConfig(path string, data []byte) error {
	var menuConfig configTypes.MenuConfig
	if err := json.Unmarshal(data, &menuConfig); err != nil {
		return fmt.Errorf("error unmarshaling menu config file %s: %w", path, err)
	}
	if bc.Menus == nil {
		bc.Menus = make(map[string]*configTypes.MenuConfig)
	}
	bc.Menus[menuConfig.Name] = &menuConfig
	return nil
}

func (bc *BotConfig) parseEncounterConfig(path string, data []byte) error {
	var encounterConfig configTypes.EncounterConfig
	if err := json.Unmarshal(data, &encounterConfig); err != nil {
		return fmt.Errorf("error unmarshaling encounter config file %s: %w", path, err)
	}
	if bc.Encounters == nil {
		bc.Encounters = make(map[string]*configTypes.EncounterConfig)
	}
	bc.Encounters[encounterConfig.Name] = &encounterConfig
	return nil
}

func (bc *BotConfig) parseConfigFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// Skip non-JSON files and directories
	if info.IsDir() || filepath.Ext(path) != ".json" {
		return nil
	}

	file, e := os.ReadFile(path)
	if e != nil {
		return fmt.Errorf("error reading config file %s: %s", path, e)
	}

	lastDirInPath := filepath.Base(filepath.Dir(path))
	switch {
	case lastDirInPath == "menus":
		if e = bc.parseMenuConfig(path, file); e != nil {
			return e
		}
	case lastDirInPath == "ultimates":
	case lastDirInPath == "savages":
	case lastDirInPath == "extremes":
		if e = bc.parseEncounterConfig(path, file); e != nil {
			return e
		}
	default:
		return fmt.Errorf("unknown config file type for file %s", path)
	}

	return nil
}

func (bc *BotConfig) Init(configDir string) {
	e := filepath.Walk(configDir, bc.parseConfigFile)
	if e != nil {
		panic(fmt.Errorf("error walking config dir: %s", e))
	}
}

func main() {
	bc := &BotConfig{}
	bc.Init("./config/data")

	fmt.Printf("Loaded BotConfig: %+v\n", bc)
}
