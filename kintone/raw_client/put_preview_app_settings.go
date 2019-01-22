package raw_client

import (
	"context"
)

type PutPreviewAppSettingsRequest struct {
	App         string `json:"app"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Theme       string `json:"theme,omitempty"`
	Revision    string `json:"revision,omitempty"`
}

type PutPreviewAppSettingsResponse struct {
	Revision string `json:"revision"`
}

func PutPreviewAppSettings(ctx context.Context, apiClient *ApiClient, req PutPreviewAppSettingsRequest) (*PutPreviewAppSettingsResponse, error) {
	apiRequest := ApiRequest{
		Method: "PUT",
		Scheme: "https",
		Path:   "/k/v1/preview/app/settings.json",
		Json:   req,
	}

	var postAppResponse PutPreviewAppSettingsResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
