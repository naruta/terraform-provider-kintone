package raw_client

import (
	"context"
)

type GetAppFormFieldsRequest struct {
	App string `json:"app"`
}

type GetAppFormFieldsResponse struct {
	Properties map[string]GetAppFormFieldsRequestProperty `json:"properties"`
	Revision   string                                     `json:"revision"`
}

type GetAppFormFieldsRequestProperty struct {
	FieldProperty
}

func GetAppFormFields(ctx context.Context, apiClient *ApiClient, req GetAppFormFieldsRequest) (*GetAppFormFieldsResponse, error) {
	apiRequest := ApiRequest{
		Method: "GET",
		Scheme: "https",
		Path:   "/k/v1/app/form/fields.json",
		Json:   req,
	}

	var GetAppFormFieldsResponse GetAppFormFieldsResponse
	if err := apiClient.Call(ctx, apiRequest, &GetAppFormFieldsResponse); err != nil {
		return nil, err
	}

	return &GetAppFormFieldsResponse, nil
}
