package starter

import (
	"context"
	"sync"

	"github.com/covrom/hex_arch_example/best_practice/internal/app/repos/user"
)

type App struct {
	us *user.Users
}

func NewApp(ust user.UserStore) *App {
	a := &App{
		us: user.NewUsers(ust),
	}
	return a
}

type APIServer interface {
	Start(us *user.Users)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs APIServer) {
	defer wg.Done()
	hs.Start(a.us)
	<-ctx.Done()
	hs.Stop()
}
