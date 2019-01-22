package kintone

import "context"

type CreateApplicationUseCaseCommand struct {
	Setting Setting
	Status  Status
	Fields  []Field
}

type CreateApplicationUseCase struct {
	apiClient ApiClient
}

func NewCreateApplicationUseCase(apiClient ApiClient) *CreateApplicationUseCase {
	return &CreateApplicationUseCase{apiClient: apiClient}
}

func (uc *CreateApplicationUseCase) Execute(ctx context.Context, cmd CreateApplicationUseCaseCommand) (AppId, Revision, error) {
	appId, revision, err := uc.apiClient.CreatePreviewApplication(ctx, cmd.Setting.Name)
	if err != nil {
		return appId, revision, err
	}

	revision, err = uc.apiClient.UpdatePreviewApplicationSettings(ctx, appId, revision, cmd.Setting)
	if err != nil {
		return appId, revision, err
	}

	revision, err = uc.apiClient.CreatePreviewApplicationFormFields(ctx, appId, revision, cmd.Fields)
	if err != nil {
		return appId, revision, err
	}

	revision, err = uc.apiClient.UpdatePreviewApplicationStatus(ctx, appId, revision, cmd.Status)
	if err != nil {
		return appId, revision, err
	}

	if err := uc.apiClient.DeployApplication(ctx, appId, revision); err != nil {
		return appId, revision, err
	}

	return appId, revision, nil
}
