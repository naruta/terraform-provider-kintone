package raw_client

import (
	"context"
)

type PutPreviewAppViewsRequestView struct {
	Index  string   `json:"index"`
	Type   string   `json:"type"`
	Name   string   `json:"name"`
	Fields []string `json:"fields"`
}

type PutPreviewAppViewsRequest struct {
	App      string                                   `json:"app"`
	Views    map[string]PutPreviewAppViewsRequestView `json:"views"`
	Revision string                                   `json:"revision,omitempty"`
}

type PutPreviewAppViewsResponse struct {
	Revision string `json:"revision"`
}

func PutPreviewAppViews(ctx context.Context, apiClient *ApiClient, req PutPreviewAppViewsRequest) (*PutPreviewAppViewsResponse, error) {
	apiRequest := ApiRequest{
		Method: "PUT",
		Scheme: "https",
		Path:   "/k/v1/preview/app/views.json",
		Json:   req,
	}

	var postAppResponse PutPreviewAppViewsResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
