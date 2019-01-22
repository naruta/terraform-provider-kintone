package raw_client

import (
	"context"
)

type PostPreviewAppDeployRequestApp struct {
	App      string `json:"app"`
	Revision string `json:"revision,omitempty"`
}

type PostPreviewAppDeployRequest struct {
	Apps   []PostPreviewAppDeployRequestApp `json:"apps"`
	Revert *bool                            `json:"revert,omitempty"`
}

type PostPreviewAppDeployResponse struct {
}

func PostPreviewAppDeploy(ctx context.Context, apiClient *ApiClient, req PostPreviewAppDeployRequest) (*PostPreviewAppDeployResponse, error) {
	apiRequest := ApiRequest{
		Method: "POST",
		Scheme: "https",
		Path:   "/k/v1/preview/app/deploy.json",
		Json:   req,
	}

	var postAppResponse PostPreviewAppDeployResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
