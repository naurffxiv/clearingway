package types

type Difficulty string

const (
	DifficultyNormal          Difficulty = "Normal"
	DifficultySavage          Difficulty = "Savage"
	DifficultyExtreme         Difficulty = "Extreme"
	DifficultyUltimate        Difficulty = "Ultimate"
	DifficultyUnreal          Difficulty = "Unreal"
	DifficultyQuantum         Difficulty = "Quantum"
	DifficultyCriterion       Difficulty = "Criterion"
	DifficultySavageCriterion Difficulty = "Savage Criterion"
	DifficultyChaotic         Difficulty = "Chaotic"
	DifficultyFields          Difficulty = "Field"
)

type RoleType string

const (
	RoleTypeCleared   RoleType = "Cleared"
	RoleTypeProg      RoleType = "Prog"
	RoleTypeRecleared RoleType = "Reclear"
	RoleTypeC4X       RoleType = "C4X"
	RoleTypeNameColor RoleType = "Name Color"
)