package methods

import (
	"context"
	"net/http"	
	"encoding/json"
)

func DecodePersistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request PersistRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}