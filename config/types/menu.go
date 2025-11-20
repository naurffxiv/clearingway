package types

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
	Style    MenuConfigButonStyle `json:"style"`
	MenuName string               `json:"menuName"`
	MenuType string               `json:"menuType"`
}

type MenuConfigRole struct {
	Name        string   `json:"name"`
	Type        RoleType `json:"type,omitempty"`
	Color       string   `json:"color,omitempty"`
	Hoist       bool     `json:"hoist,omitempty"`
	Mention     bool     `json:"mention,omitempty"`
	Description string   `json:"description,omitempty"`
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
	MultiSelect  bool               `json:"multiSelect,omitempty"`
	RequireClear bool               `json:"requireClear,omitempty"`
	Difficulties []Difficulty       `json:"difficulties,omitempty"`
	RoleType     []string           `json:"roleType,omitempty"`
	Roles        []MenuConfigRole   `json:"roles,omitempty"`
}
