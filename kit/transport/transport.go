package transport

import(
	"zkgo/kit/endpoints"
	"github.com/go-kit/kit/log"
	"net/http"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
	"encoding/json"
	"context"
)

func NewHttpHandler(endpoints endpoint.Endpoints, logger log.Logger) http.Handler{
	router := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
	}
	router.Methods("POST").Path("/zkgo").Handler(
		httptransport.NewServer(
			endpoints.RefreshKeyEndpoint,
			decodeHTTPRefreshKeyRequest,
			encodeHTTPRefreshKeyResponse,
			options...),
	)
	return router
}

func decodeHTTPRefreshKeyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.RefreshKeyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encodeHTTPRefreshKeyResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}