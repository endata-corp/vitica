package vitica

import (
	"time"
	"vitica/_vendor/src/github.com/gocraft/web"
)

func LoggerMiddleware(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	startTime := time.Now()

	next(rw, req)

	duration := time.Since(startTime).Nanoseconds()
	var durationUnits string
	switch {
	case duration > 2000000:
		durationUnits = "ms"
		duration /= 1000000
	case duration > 1000:
		durationUnits = "μs"
		duration /= 1000
	default:
		durationUnits = "ns"
	}

	Info("[%d %s] %d '%s'\n", duration, durationUnits, rw.StatusCode(), req.URL.Path)
}
