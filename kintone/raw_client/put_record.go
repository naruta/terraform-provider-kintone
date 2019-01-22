package raw_client

import (
	"context"
)

type PutRecordRequestRecord struct {
	Value string `json:"value"`
}

type PutRecordRequest struct {
	App    string                            `json:"app"`
	Id     string                            `json:"id"`
	Record map[string]PutRecordRequestRecord `json:"record,omitempty"`
}

type PutRecordResponse struct {
	Revision string `json:"revision"`
}

func PutRecord(ctx context.Context, apiClient *ApiClient, req PutRecordRequest) (*PutRecordResponse, error) {
	apiRequest := ApiRequest{
		Method: "PUT",
		Scheme: "https",
		Path:   "/k/v1/record.json",
		Json:   req,
	}

	var postAppResponse PutRecordResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
