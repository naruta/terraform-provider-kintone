package raw_client

import (
	"context"
)

type DeleteRecordsRequestRecord struct {
	Value string `json:"value"`
}

type DeleteRecordsRequest struct {
	App       string   `json:"app"`
	Ids       []string `json:"ids"`
	Revisions []string `json:"revisions,omitempty"`
}

type DeleteRecordsResponse struct {
}

func DeleteRecords(ctx context.Context, apiClient *ApiClient, req DeleteRecordsRequest) (*DeleteRecordsResponse, error) {
	apiRequest := ApiRequest{
		Method: "DELETE",
		Scheme: "https",
		Path:   "/k/v1/records.json",
		Json:   req,
	}

	var postAppResponse DeleteRecordsResponse
	if err := apiClient.Call(ctx, apiRequest, &postAppResponse); err != nil {
		return nil, err
	}

	return &postAppResponse, nil
}
