package raw_client

import (
	"context"
)

type GetPreviewAppDeployRequest struct {
	Apps []string `json:"apps"`
}

type GetPreviewAppDeployResponse struct {
	Apps []GetPreviewAppDeployResponseApp `json:"apps"`
}
type GetPreviewAppDeployResponseApp struct {
	App    string `json:"app"`
	Status string `json:"status"`
}

func GetPreviewAppDeploy(ctx context.Context, apiClient *ApiClient, req GetPreviewAppDeployRequest) (*GetPreviewAppDeployResponse, error) {
	apiRequest := ApiRequest{
		Method: "GET",
		Scheme: "https",
		Path:   "/k/v1/preview/app/deploy.json",
		Json:   req,
	}

	var GetAppResponse GetPreviewAppDeployResponse
	if err := apiClient.Call(ctx, apiRequest, &GetAppResponse); err != nil {
		return nil, err
	}

	return &GetAppResponse, nil
}
