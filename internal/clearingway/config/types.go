package config

type Difficulty string

const (
	DifficultySavage   Difficulty = "Savage"
	DifficultyExtreme  Difficulty = "Extreme"
	DifficultyUltimate Difficulty = "Ultimate"
)

type RoleType string

const (
	RoleTypeCleared   RoleType = "Cleared"
	RoleTypeProg      RoleType = "Prog"
	RoleTypeRecleared RoleType = "Reclear"
	RoleTypeC4X       RoleType = "C4X"
	RoleTypeNameColor RoleType = "Name Color"
)

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

type MenuConfigType string

const (
	MenuConfigTypeMain      MenuConfigType = "menuMain"
	MenuConfigTypeVerify    MenuConfigType = "menuVerify"
	MenuConfigTypeRemove    MenuConfigType = "menuRemove"
	MenuConfigTypeEncounter MenuConfigType = "menuEncounter"
)

type MenuConfigField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type MenuConfigButonStyle int

const (
	MenuConfigButtonStylePrimary   MenuConfigButonStyle = 1
	MenuConfigButtonStyleSecondary MenuConfigButonStyle = 2
	MenuConfigButtonStyleSuccess   MenuConfigButonStyle = 3
	MenuConfigButtonStyleDanger    MenuConfigButonStyle = 4
)

type MenuConfigButton struct {
	Label    string               `json:"label"`
	MenuName string               `json:"menuName"`
	MenuType string               `json:"menuType"`
	Style    MenuConfigButonStyle `json:"style"`
}

type MenuConfigRole struct {
	Name        string   `json:"name"`
	Type        RoleType `json:"type,omitempty"`
	Color       string   `json:"color,omitempty"`
	Description string   `json:"description,omitempty"`
	Hoist       bool     `json:"hoist,omitempty"`
	Mention     bool     `json:"mention,omitempty"`
}

type MenuConfig struct {
	Name         string             `json:"name"`
	Type         MenuConfigType     `json:"type"`
	Title        string             `json:"title,omitempty"`
	Description  string             `json:"description,omitempty"`
	ThumbnailURL string             `json:"thumbnailUrl,omitempty"`
	ImageURL     string             `json:"imageUrl,omitempty"`
	Fields       []MenuConfigField  `json:"fields,omitempty"`
	Buttons      []MenuConfigButton `json:"buttons,omitempty"`
	Difficulties []Difficulty       `json:"difficulties,omitempty"`
	RoleType     []string           `json:"roleType,omitempty"`
	Roles        []MenuConfigRole   `json:"roles,omitempty"`
	MultiSelect  bool               `json:"multiSelect,omitempty"`
	RequireClear bool               `json:"requireClear,omitempty"`
}
