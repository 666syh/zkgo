package endpoint

import(
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"time"
	"context"
)

func LoggingMiddleware(logger log.Logger) endpoint.Middleware{
	return func(next endpoint.Endpoint) endpoint.Endpoint{
		return func(ctx context.Context, request interface{}) (res interface{}, err error){
			defer func(begin time.Time){
				if err!=nil{
					logger.Log("time", begin, "Error", err)
				}
			}(time.Now())
			return next(ctx, request)
		}
	}
}