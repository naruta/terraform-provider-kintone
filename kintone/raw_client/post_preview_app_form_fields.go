package raw_client

import "context"

type PostPreviewAppFormFieldsRequestProperty struct {
	Code  string `json:"code"`
	Label string `json:"label"`
	Type  string `json:"type"`
}

type PostPreviewAppFormFieldsRequest struct {
	App        string                                             `json:"app"`
	Properties map[string]PostPreviewAppFormFieldsRequestProperty `json:"properties"`
	Revision   string                                             `json:"revision,omitempty"`
}

type PostPreviewAppFormFieldsResponse struct {
	Revision string `json:"revision"`
}

func PostPreviewAppFormFields(ctx context.Context, apiClient *ApiClient, req PostPreviewAppFormFieldsRequest) (*PostPreviewAppFormFieldsResponse, error) {
	apiRequest := ApiRequest{
		Method: "POST",
		Scheme: "https",
		Path:   "/k/v1/preview/app/form/fields.json",
		Json:   req,
	}

	var postAppResponse PostPreviewAppFormFieldsResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
