package userrepo

import (
	"context"
	"fmt"

	"github.com/covrom/hex_arch_example/best_practice/internal/entities/user"

	"github.com/google/uuid"
)

// only needed here
type UserStore interface {
	Create(ctx context.Context, u user.User) (*uuid.UUID, error)
	Read(ctx context.Context, uid uuid.UUID) (*user.User, error)
	Delete(ctx context.Context, uid uuid.UUID) error
	SearchUsers(ctx context.Context, s string) (chan user.User, error)
}

type Users struct {
	ustore UserStore
}

func NewUsers(ustore UserStore) *Users {
	return &Users{
		ustore: ustore,
	}
}

func (us *Users) Create(ctx context.Context, u user.User) (*user.User, error) {
	u.ID = uuid.New()
	id, err := us.ustore.Create(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}
	u.ID = *id
	return &u, nil
}

func (us *Users) Read(ctx context.Context, uid uuid.UUID) (*user.User, error) {
	u, err := us.ustore.Read(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("read user error: %w", err)
	}
	return u, nil
}

func (us *Users) Delete(ctx context.Context, uid uuid.UUID) (*user.User, error) {
	u, err := us.ustore.Read(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("search user error: %w", err)
	}
	return u, us.ustore.Delete(ctx, uid)
}

func (us *Users) SearchUsers(ctx context.Context, s string) (chan user.User, error) {
	// FIXME: здесь нужно использвоать паттерн Unit of Work
	// бизнес-транзакция
	chin, err := us.ustore.SearchUsers(ctx, s)
	if err != nil {
		return nil, err
	}
	chout := make(chan user.User, 100)
	go func() {
		defer close(chout)
		for {
			select {
			case <-ctx.Done():
				return
			case u, ok := <-chin:
				if !ok {
					return
				}
				u.Permissions = 0755
				chout <- u
			}
		}
	}()
	return chout, nil
}
