package dto

type (
	CreateTournamentRequest struct {
		Game        uint8 `json:"game" validate:"required"`
		Organizer   string
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Location    string `json:"location" validate:"required"`
		StartDate   string `json:"startDate" validate:"required"`
		EndDate     string `json:"endDate" validate:"required"`
		Contact     string `json:"contact" validate:"required"`
		Prize       uint64 `json:"prize" validate:"required"`
		MaxTeam     uint16 `json:"maxTeam" validate:"required"`
	}
)
