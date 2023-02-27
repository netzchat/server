package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/netzchat/server/apis/core/v1/corev1connect"
	"github.com/netzchat/server/apis/test/v1/testv1connect"
	"github.com/netzchat/server/ent"
	"github.com/netzchat/server/services/channel"
	"github.com/netzchat/server/services/test"
)

func run(ctx context.Context) error {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return fmt.Errorf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()

	if err := client.Schema.Create(ctx); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	http.Handle(testv1connect.NewTestServiceHandler(test.New()))
	http.Handle(corev1connect.NewChannelServiceHandler(channel.New(client)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return fmt.Errorf("error running server: %v", err)
	}

	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("error running server: %s", err)
	}
}
