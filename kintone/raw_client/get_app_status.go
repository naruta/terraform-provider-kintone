package raw_client

import (
	"context"
)

type GetAppStatusRequest struct {
	App string `json:"app"`
}

type GetAppStatusResponse struct {
	Enable   bool                                `json:"enable"`
	States   map[string]GetAppStatusRequestState `json:"states"`
	Actions  []GetAppStatusRequestAction         `json:"actions"`
	Revision string                              `json:"revision"`
}

type GetAppStatusRequestState struct {
	Name  string `json:"name"`
	Index string `json:"index"`
}

type GetAppStatusRequestAction struct {
	Name string `json:"name"`
	From string `json:"from"`
	To   string `json:"to"`
}

func GetAppStatus(ctx context.Context, apiClient *ApiClient, req GetAppStatusRequest) (*GetAppStatusResponse, error) {
	apiRequest := ApiRequest{
		Method: "GET",
		Scheme: "https",
		Path:   "/k/v1/app/status.json",
		Json:   req,
	}

	var GetAppStatusResponse GetAppStatusResponse
	if err := apiClient.Call(ctx, apiRequest, &GetAppStatusResponse); err != nil {
		return nil, err
	}

	return &GetAppStatusResponse, nil
}
