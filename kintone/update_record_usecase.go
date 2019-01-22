package kintone

import "context"

type UpdateRecordUseCaseCommand struct {
	AppId  AppId
	Record Record
}

type UpdateRecordUseCase struct {
	apiClient ApiClient
}

func NewUpdateRecordUseCase(apiClient ApiClient) *UpdateRecordUseCase {
	return &UpdateRecordUseCase{apiClient: apiClient}
}

func (uc *UpdateRecordUseCase) Execute(ctx context.Context, cmd UpdateRecordUseCaseCommand) (RecordRevision, error) {
	return uc.apiClient.UpdateRecord(ctx, cmd.AppId, cmd.Record)
}
