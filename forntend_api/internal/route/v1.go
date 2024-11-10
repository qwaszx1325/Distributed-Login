package route

import "go.uber.org/fx"

type RouteV1 struct {
}

var _ Route = (*RouteV1)(nil)

func NewRouteV1Set() fx.Option {

}
