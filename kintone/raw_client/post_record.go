package raw_client

import (
	"context"
)

type PostRecordRequestRecord struct {
	Value string `json:"value"`
}

type PostRecordRequest struct {
	App    string                             `json:"app"`
	Record map[string]PostRecordRequestRecord `json:"record,omitempty"`
}

type PostRecordResponse struct {
	Id       string `json:"id"`
	Revision string `json:"revision"`
}

func PostRecord(ctx context.Context, apiClient *ApiClient, req PostRecordRequest) (*PostRecordResponse, error) {
	apiRequest := ApiRequest{
		Method: "POST",
		Scheme: "https",
		Path:   "/k/v1/record.json",
		Json:   req,
	}

	var postAppResponse PostRecordResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
