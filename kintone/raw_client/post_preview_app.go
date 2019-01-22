package raw_client

import (
	"context"
)

type PostPreviewAppRequest struct {
	Name   string `json:"name"`
	Space  *int   `json:"space,omitempty"`
	Thread *int   `json:"thread,omitempty"`
}

type PostPreviewAppResponse struct {
	App      string `json:"app"`
	Revision string `json:"revision"`
}

func PostPreviewApp(ctx context.Context, apiClient *ApiClient, req PostPreviewAppRequest) (*PostPreviewAppResponse, error) {
	apiRequest := ApiRequest{
		Method: "POST",
		Scheme: "https",
		Path:   "/k/v1/preview/app.json",
		Json:   req,
	}

	var postAppResponse PostPreviewAppResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
