package starter

import (
	"context"
	"sync"

	"github.com/covrom/hex_arch_example/best_practice/internal/logic/app/repos/userrepo"
)

type App struct {
	us *userrepo.Users
}

func NewApp(ust userrepo.UserStore) *App {
	a := &App{
		us: userrepo.NewUsers(ust),
	}
	return a
}

type APIServer interface {
	Start(us *userrepo.Users)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs APIServer) {
	defer wg.Done()
	hs.Start(a.us)
	<-ctx.Done()
	hs.Stop()
}
