package seeder

import (
	"context"
	"ecom-mono-go/domain/service"
	"ecom-mono-go/domain/types"
	"ecom-mono-go/infrastructure"
)

type Seeder struct {
	us service.UserService
	env *infrastructure.Env
}

func NewSeeder(us service.UserService, env *infrastructure.Env) *Seeder{
	return &Seeder{
		us: us,
		env: env,
	}
}

func (s *Seeder) Seed() {
	// TODO log errors
	ctx := context.Background()
	s.seedDummyUser(ctx)
	s.seedDummyVendor(ctx)
	s.seedSuperAdmin(ctx)
}

func (s *Seeder) seedDummyUser(ctx context.Context) error {
	id := types.NewID()
	_,err := s.us.CreateUser(
		ctx,
		&types.User{
			ID: id,
			Username: s.env.DUMMY_USERNAME,
			Email: s.env.DUMMY_USER_EMAIL,
			EmailVerified: true,
			Password: s.env.DUMMY_PASSWORD,
			Base: &types.Base{
				CreatedByID: id,
				UpdatedByID: id,
			},
		},
	)
	return err
}

func (s *Seeder) seedDummyVendor(ctx context.Context) error {
	id := types.NewID()
	_,err := s.us.CreateUser(
		ctx,
		&types.User{
			ID: id,
			Username: s.env.DUMMY_USERNAME,
			Email: s.env.DUMMY_VENDOR_EMAIL,
			EmailVerified: true,
			Password: s.env.DUMMY_PASSWORD,
			Role: types.ROLE_VENDOR,
			Base: &types.Base{
				CreatedByID: id,
				UpdatedByID: id,
			},
		},
	)
	return err
}

func (s *Seeder) seedSuperAdmin(ctx context.Context) error {
	id := types.NewID()
	_,err := s.us.CreateUser(
		ctx,
		&types.User{
			ID: id,
			Username: s.env.SUPER_ADMIN_USERNAME,
			Email: s.env.SUPER_ADMIN_EMAIL,
			EmailVerified: true,
			Password: s.env.SUPER_ADMIN_PASSWORD,
			Role: types.ROLE_SUPER_ADMIN,
			Base: &types.Base{
				CreatedByID: id,
				UpdatedByID: id,
			},
		},
	)
	return err
}