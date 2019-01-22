package raw_client

import (
	"context"
)

type GetAppSettingsRequest struct {
	App string `json:"app"`
}

type GetAppSettingsResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Theme       string `json:"Theme"`
}

func GetAppSettings(ctx context.Context, apiClient *ApiClient, req GetAppSettingsRequest) (*GetAppSettingsResponse, error) {
	apiRequest := ApiRequest{
		Method: "GET",
		Scheme: "https",
		Path:   "/k/v1/app/settings.json",
		Json:   req,
	}

	var GetAppSettingsResponse GetAppSettingsResponse
	if err := apiClient.Call(ctx, apiRequest, &GetAppSettingsResponse); err != nil {
		return nil, err
	}

	return &GetAppSettingsResponse, nil
}
