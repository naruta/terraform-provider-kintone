package raw_client

import (
	"context"
)

type GetPreviewAppViewsRequest struct {
	App string `json:"app"`
}

type GetPreviewAppViewsResponse struct {
	Views map[string]GetPreviewAppViewsResponseView `json:"views"`
}
type GetPreviewAppViewsResponseView struct {
	Type        string   `json:"type"`
	BuiltinType string   `json:"builtinType"`
	Name        string   `json:"name"`
	Id          string   `json:"id"`
	Fields      []string `json:"fields"`
	Date        string   `json:"date"`
	Title       string   `json:"title"`
	Html        string   `json:"html"`
	Pager       bool     `json:"pager"`
	Device      string   `json:"device"`
	FilterCond  string   `json:"filterCond"`
	Sort        string   `json:"sort"`
	Index       string   `json:"index"`
	Revision    string   `json:"revision"`
}

func GetPreviewAppViews(ctx context.Context, apiClient *ApiClient, req GetPreviewAppViewsRequest) (*GetPreviewAppViewsResponse, error) {
	apiRequest := ApiRequest{
		Method: "GET",
		Scheme: "https",
		Path:   "/k/v1/preview/app/views.json",
		Json:   req,
	}

	var GetAppResponse GetPreviewAppViewsResponse
	if err := apiClient.Call(ctx, apiRequest, &GetAppResponse); err != nil {
		return nil, err
	}

	return &GetAppResponse, nil
}
