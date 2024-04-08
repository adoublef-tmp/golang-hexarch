package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	iamRedis "github.com/roku-on-it/golang-search/internal/iam/redis"
	"github.com/roku-on-it/golang-search/net/http"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background(), os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error occurred: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string) error {
	var (
		fs   = flag.NewFlagSet("", flag.ContinueOnError)
		port = fs.Int("port", 8080, "http listening port")
		dsn  = fs.String("dsn", ":redis", "redis address")
	)
	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("parsing args error: %w", err)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	rc := redis.NewClient(&redis.Options{Addr: *dsn}) // no_password,default_db not recommended for production

	app := http.App(&iamRedis.DB{Conn: rc})

	g.Go(func() error {
		return app.Listen(fmt.Sprintf(":%d", *port))
	})

	g.Go(func() error {
		<-ctx.Done()
		ctx, timeout := context.WithTimeout(context.Background(), time.Second*5)
		defer timeout()
		return app.ShutdownWithContext(ctx)
	})

	return g.Wait()
}
