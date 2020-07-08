package framework

type route struct {
	app *app
	route string
}

func (r *route) CreateHandler() *handler {
	return newHandler(r.app,r.route)
}

func (r *route) SubRoute(path string) *route {
	return &route{app: r.app,route: r.route+path}
}