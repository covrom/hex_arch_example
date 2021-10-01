package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/covrom/hex_arch_example/best_practice/internal/api/handler"
	"github.com/covrom/hex_arch_example/best_practice/internal/api/server"
	"github.com/covrom/hex_arch_example/best_practice/internal/app/repos/user"
	"github.com/covrom/hex_arch_example/best_practice/internal/app/starter"
	"github.com/covrom/hex_arch_example/best_practice/internal/db/mem/usermemstore"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	ust := usermemstore.NewUsers()
	a := starter.NewApp(ust)
	us := user.NewUsers(ust)
	h := handler.NewRouter(us)
	srv := server.NewServer(":8000", h)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
