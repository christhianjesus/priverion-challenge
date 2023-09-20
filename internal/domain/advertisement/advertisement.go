package advertisement

import "time"

type Advertisement interface {
	ID() string
	UserID() string
	Title() string
	Body() string
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

func ValidateAdvertisement(advertisement Advertisement) error {
	return nil
}
