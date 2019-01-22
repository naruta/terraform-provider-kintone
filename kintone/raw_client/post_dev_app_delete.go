package raw_client

import (
	"context"
)

type PostDevAppDeleteRequest struct {
	App string `json:"app"`
	/*RequestToken string `json:"__REQUEST_TOKEN__"`*/
}

type PostDevAppDeleteResponse struct {
}

func PostDevAppDelete(ctx context.Context, apiClient *ApiClient, req PostDevAppDeleteRequest) (*PostDevAppDeleteResponse, error) {
	apiRequest := ApiRequest{
		Method: "POST",
		Scheme: "https",
		Path:   "/k/api/dev/app/delete.json",
		Json:   req,
	}

	var postAppResponse PostDevAppDeleteResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
