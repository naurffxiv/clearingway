package types

type EncounterRole struct {
	Name  string   `json:"name"`
	Type  RoleType `json:"type"`
	Color string   `json:"color"`
}

type EncounterConfig struct {
	IDs        []int           `json:"ids"`
	Name       string          `json:"name"`
	Difficulty Difficulty      `json:"difficulty"`
	Roles      []EncounterRole `json:"roles"`
}
