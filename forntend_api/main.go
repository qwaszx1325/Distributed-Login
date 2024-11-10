package main

import (
	"net/http"

	"example.com/simple-login/forntend_api/internal/route"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		route.NewRouteV1Set(),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
