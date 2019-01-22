package kintone

import "context"

type UpdateApplicationUseCaseCommand struct {
	AppId            AppId
	Revision         Revision
	Setting          Setting
	Status           Status
	CreateFields     []Field
	UpdateFields     []Field
	DeleteFieldCodes []FieldCode
}

type UpdateApplicationUseCase struct {
	apiClient ApiClient
}

func NewUpdateApplicationUseCase(apiClient ApiClient) *UpdateApplicationUseCase {
	return &UpdateApplicationUseCase{apiClient: apiClient}
}

func (uc *UpdateApplicationUseCase) Execute(ctx context.Context, cmd UpdateApplicationUseCaseCommand) (Revision, error) {
	appId := cmd.AppId

	revision, err := uc.apiClient.UpdatePreviewApplicationSettings(ctx, appId, cmd.Revision, cmd.Setting)
	if err != nil {
		return revision, err
	}

	if len(cmd.CreateFields) > 0 {
		revision, err = uc.apiClient.CreatePreviewApplicationFormFields(ctx, appId, revision, cmd.CreateFields)
		if err != nil {
			return revision, err
		}
	}

	if len(cmd.UpdateFields) > 0 {
		revision, err = uc.apiClient.UpdatePreviewApplicationFormFields(ctx, appId, revision, cmd.UpdateFields)
		if err != nil {
			return revision, err
		}
	}

	if len(cmd.DeleteFieldCodes) > 0 {
		revision, err = uc.apiClient.DeletePreviewApplicationFormFields(ctx, appId, revision, cmd.DeleteFieldCodes)
		if err != nil {
			return revision, err
		}
	}

	revision, err = uc.apiClient.UpdatePreviewApplicationStatus(ctx, appId, revision, cmd.Status)
	if err != nil {
		return revision, err
	}

	if err := uc.apiClient.DeployApplication(ctx, appId, revision); err != nil {
		return revision, err
	}

	return revision, nil
}
