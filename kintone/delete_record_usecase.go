package kintone

import "context"

type DeleteRecordUseCaseCommand struct {
	AppId AppId
	Id    RecordId
}

type DeleteRecordUseCase struct {
	apiClient ApiClient
}

func NewDeleteRecordUseCase(apiClient ApiClient) *DeleteRecordUseCase {
	return &DeleteRecordUseCase{apiClient: apiClient}
}

func (uc *DeleteRecordUseCase) Execute(ctx context.Context, cmd DeleteRecordUseCaseCommand) error {
	return uc.apiClient.DeleteRecords(ctx, cmd.AppId, []RecordId{cmd.Id})
}
