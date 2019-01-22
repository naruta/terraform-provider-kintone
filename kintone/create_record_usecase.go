package kintone

import "context"

type CreateRecordUseCaseCommand struct {
	AppId  AppId
	Record Record
}

type CreateRecordUseCase struct {
	apiClient ApiClient
}

func NewCreateRecordUseCase(apiClient ApiClient) *CreateRecordUseCase {
	return &CreateRecordUseCase{apiClient: apiClient}
}

func (uc *CreateRecordUseCase) Execute(ctx context.Context, cmd CreateRecordUseCaseCommand) (RecordId, RecordRevision, error) {
	return uc.apiClient.CreateRecord(ctx, cmd.AppId, cmd.Record)
}
