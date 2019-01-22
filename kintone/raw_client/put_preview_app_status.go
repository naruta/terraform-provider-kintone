package raw_client

import (
	"context"
)

type PutPreviewAppStatusRequestState struct {
	Name  string `json:"name"`
	Index string `json:"index"`
}

type PutPreviewAppStatusRequestAction struct {
	Name string `json:"name"`
	From string `json:"from"`
	To   string `json:"to"`
}

type PutPreviewAppStatusRequest struct {
	App      string                                     `json:"app"`
	Enable   bool                                       `json:"enable"`
	States   map[string]PutPreviewAppStatusRequestState `json:"states,omitempty"`
	Actions  []PutPreviewAppStatusRequestAction         `json:"actions,omitempty"`
	Revision string                                     `json:"revision,omitempty"`
}

type PutPreviewAppStatusResponse struct {
	Revision string `json:"revision"`
}

func PutPreviewAppStatus(ctx context.Context, apiClient *ApiClient, req PutPreviewAppStatusRequest) (*PutPreviewAppStatusResponse, error) {
	apiRequest := ApiRequest{
		Method: "PUT",
		Scheme: "https",
		Path:   "/k/v1/preview/app/status.json",
		Json:   req,
	}

	var postAppResponse PutPreviewAppStatusResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
