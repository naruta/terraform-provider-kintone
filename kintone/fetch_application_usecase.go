package kintone

import "context"

type FetchApplicationUseCaseQuery struct {
	AppId AppId
}

type FetchApplicationUseCase struct {
	apiClient ApiClient
}

func NewFetchApplicationUseCase(apiClient ApiClient) *FetchApplicationUseCase {
	return &FetchApplicationUseCase{apiClient: apiClient}
}

func (uc *FetchApplicationUseCase) Execute(ctx context.Context, cmd FetchApplicationUseCaseQuery) (Application, error) {
	app, err := uc.apiClient.FetchApplication(ctx, cmd.AppId)
	if err != nil {
		return Application{}, err
	}

	var handleFields []Field
	for _, f := range app.Fields {
		handleFields = append(handleFields, f)
	}
	app.Fields = handleFields
	return app, nil
}
