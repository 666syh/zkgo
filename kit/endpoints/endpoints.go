package endpoint

import(
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"zkgo/kit/service"
	"context"
)

type Endpoints struct{
	RefreshKeyEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc service.Service, logger log.Logger) Endpoints{
	refreshkeyEndpoint := MakeRefreshKeyEndpoint(svc)
	refreshkeyEndpoint = LoggingMiddleware(logger)(refreshkeyEndpoint)
	return Endpoints{
		refreshkeyEndpoint,
	}
}

func MakeRefreshKeyEndpoint(svc service.Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(RefreshKeyRequest)
		err := svc.RefreshKey(ctx, req.Target)
		if err != nil{
			return RefreshKeyResponse{err}, err
		} else{
			return RefreshKeyResponse{nil}, nil
		}
	}
}

type RefreshKeyRequest struct{
	Target interface{} `json:"target"`
}

type RefreshKeyResponse struct{
	Err error `json:"err"`
}