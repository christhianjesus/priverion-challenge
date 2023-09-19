package advertisement

import (
	"context"
)

type AdvertisementService interface {
	GetAll(ctx context.Context) ([]Advertisement, error)
	Create(ctx context.Context, advertisement Advertisement) error
	GetOne(ctx context.Context, advertisementID string) (Advertisement, error)
	Update(ctx context.Context, advertisement Advertisement) error
	Delete(ctx context.Context, advertisementID string) error
}
