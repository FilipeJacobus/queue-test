package methods

import (
	"container/list"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakePersistEndpoint(svc StringService, queue *list.List) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(PersistRequest)
		v, err := svc.Persist(req, queue)
		if err != nil {
			return PersistResponse{v, err.Error()}, nil
		}
		return PersistResponse{v, ""}, nil
	}
}
