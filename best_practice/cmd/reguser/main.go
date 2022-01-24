package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/covrom/hex_arch_example/best_practice/internal/infrastructure/api/handler"
	"github.com/covrom/hex_arch_example/best_practice/internal/infrastructure/api/server"
	"github.com/covrom/hex_arch_example/best_practice/internal/infrastructure/db/mem/usermemstore"
	"github.com/covrom/hex_arch_example/best_practice/internal/logic/app/repos/userrepo"
	"github.com/covrom/hex_arch_example/best_practice/internal/logic/app/starter"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	ust := usermemstore.NewUsers()
	a := starter.NewApp(ust)
	us := userrepo.NewUsers(ust)
	h := handler.NewRouter(us)
	srv := server.NewServer(":8000", h)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
