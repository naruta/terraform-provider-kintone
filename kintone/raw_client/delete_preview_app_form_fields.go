package raw_client

import "context"

type DeletePreviewAppFormFieldsRequest struct {
	App      string   `json:"app"`
	Fields   []string `json:"fields"`
	Revision string   `json:"revision,omitempty"`
}

type DeletePreviewAppFormFieldsResponse struct {
	Revision string `json:"revision"`
}

func DeletePreviewAppFormFields(ctx context.Context, apiClient *ApiClient, req DeletePreviewAppFormFieldsRequest) (*DeletePreviewAppFormFieldsResponse, error) {
	apiRequest := ApiRequest{
		Method: "DELETE",
		Scheme: "https",
		Path:   "/k/v1/preview/app/form/fields.json",
		Json:   req,
	}

	var postAppResponse DeletePreviewAppFormFieldsResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
