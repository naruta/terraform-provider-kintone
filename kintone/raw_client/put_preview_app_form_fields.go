package raw_client

import "context"

type PutPreviewAppFormFieldsRequestProperty struct {
	FieldProperty
}

type PutPreviewAppFormFieldsRequest struct {
	App        string                                            `json:"app"`
	Properties map[string]PutPreviewAppFormFieldsRequestProperty `json:"properties"`
	Revision   string                                            `json:"revision,omitempty"`
}

type PutPreviewAppFormFieldsResponse struct {
	Revision string `json:"revision"`
}

func PutPreviewAppFormFields(ctx context.Context, apiClient *ApiClient, req PutPreviewAppFormFieldsRequest) (*PutPreviewAppFormFieldsResponse, error) {
	apiRequest := ApiRequest{
		Method: "PUT",
		Scheme: "https",
		Path:   "/k/v1/preview/app/form/fields.json",
		Json:   req,
	}

	var postAppResponse PutPreviewAppFormFieldsResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
