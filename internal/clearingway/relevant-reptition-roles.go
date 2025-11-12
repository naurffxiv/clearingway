package clearingway

import (
	"fmt"
)

func RelevantRepetitionRoles(encs *Encounters) *Roles {
	roles := &Roles{Roles: []*Role{
		{
			Name:        "Please Do Other Content",
			Color:       0xcf142b,
			Uncomfy:     true,
			Type:        CompleteRole,
			Description: "Cleared any relevant encounter at least 100 times.",
			ShouldApply: func(opts *ShouldApplyOpts) (bool, string) {
				for _, encounter := range opts.Encounters.Encounters {
					clears := 0

					for _, encounterId := range encounter.Ids {
						ranking, ok := opts.Rankings.Rankings[encounterId]

						if !ok {
							continue
						}
						if !ranking.Cleared() {
							continue
						}

						clears = clears + ranking.TotalKills
					}

					if clears >= 100 {
						return true,
							fmt.Sprintf(
								"Cleared `%v` at least **100** times (**%v** total).",
								encounter.Name,
								clears,
							)
					}
				}

				return false, "Did not clear any encounter at least 100 times."
			},
		},
	}}

	for _, enc := range encs.Encounters {
		if enc.TotalWeaponsAvailable == 0 {
			continue
		}

		roles.Roles = append(roles.Roles, &Role{
			Name:        "Limbo",
			Color:       0x808080,
			Uncomfy:     true,
			Type:        LimboRole,
			Encounter:   enc,
			Description: "Cleared " + enc.Name + "... but only once.",
			ShouldApply: func(opts *ShouldApplyOpts) (bool, string) {
				for _, encounter := range opts.Encounters.Encounters {
					if encounter.Name != enc.Name {
						continue
					}
					clears := 0

					for _, encounterId := range encounter.Ids {
						ranking, ok := opts.Rankings.Rankings[encounterId]

						if !ok {
							continue
						}
						if !ranking.Cleared() {
							continue
						}

						clears = clears + ranking.TotalKills
					}

					if clears == 1 {
						return true,
							fmt.Sprintf(
								"Cleared `%s`... but only **once**.\nUse `/uncomfy` if you don't want this role.",
								enc.Name,
							)
					}
				}

				return false, fmt.Sprintf("Cleared `%s` more than **once**.", enc.Name)
			},
		})

		roles.Roles = append(roles.Roles, &Role{
			Name:        "Complete",
			Color:       0xffde00,
			Type:        CompleteRole,
			Encounter:   enc,
			Description: "Cleared " + enc.Name + " at least " + enc.CompleteNumber() + " times.",
			ShouldApply: func(opts *ShouldApplyOpts) (bool, string) {
				for _, encounter := range opts.Encounters.Encounters {
					if encounter.Name != enc.Name {
						continue
					}
					clears := 0

					for _, encounterId := range encounter.Ids {
						ranking, ok := opts.Rankings.Rankings[encounterId]

						if !ok {
							continue
						}
						if !ranking.Cleared() {
							continue
						}

						clears = clears + ranking.TotalKills
					}

					if clears >= enc.TotalWeaponsAvailable {
						return true,
							fmt.Sprintf(
								"Cleared `%s` at least **%s** times.",
								enc.Name, enc.CompleteNumber(),
							)
					}
				}

				return false, fmt.Sprintf("Has not cleared `%s` at least **%s** times.", enc.Name, enc.CompleteNumber())
			},
		})
	}

	return roles
}
