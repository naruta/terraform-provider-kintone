package kintone

import "context"

type DeleteApplicationUseCaseCommand struct {
	AppId AppId
}

type DeleteApplicationUseCase struct {
	apiClient ApiClient
}

func NewDeleteApplicationUseCase(apiClient ApiClient) *DeleteApplicationUseCase {
	return &DeleteApplicationUseCase{apiClient: apiClient}
}

func (uc *DeleteApplicationUseCase) Execute(ctx context.Context, cmd DeleteApplicationUseCaseCommand) error {
	return uc.apiClient.DeleteApplication(ctx, cmd.AppId)
}
