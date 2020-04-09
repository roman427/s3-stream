package sub

import "github.com/valyala/fasthttp"

// routes assigns all routes to router.
func (e *env) routes() {
	e.router = func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			e.handleRoot(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
}
