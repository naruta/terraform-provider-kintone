package kintone

import (
	"context"
)

type CreateApplicationUseCaseCommand struct {
	Setting Setting
	Status  Status
	Fields  []Field
	Views   []View
}

type CreateApplicationUseCase struct {
	apiClient ApiClient
}

func NewCreateApplicationUseCase(apiClient ApiClient) *CreateApplicationUseCase {
	return &CreateApplicationUseCase{apiClient: apiClient}
}

func filterBuiltinView(views []View) []View {
	var builtinViews []View
	for _, v := range views {
		if len(v.BuiltinType) > 0 {
			builtinViews = append(builtinViews, v)
		}
	}
	return builtinViews
}

func containsView(views []View, view View) bool {
	for _, v := range views {
		if v.Name == view.Name {
			return true
		}
	}
	return false
}

func (uc *CreateApplicationUseCase) mergeBuiltinView(ctx context.Context, appId AppId, views []View) ([]View, error) {
	views, err := uc.apiClient.FetchPreviewApplicationViews(ctx, appId)
	if err != nil {
		return []View{}, err
	}
	builtinViews := filterBuiltinView(views)

	var mergedViews = append([]View{}, views...)
	for _, builtinView := range builtinViews {
		if containsView(views, builtinView) {
			builtinView.Index = "1000"
			mergedViews = append(mergedViews, builtinView)
		}
	}
	return mergedViews, nil
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

	newViews, err := uc.mergeBuiltinView(ctx, appId, cmd.Views)
	if err != nil {
		return appId, revision, err
	}

	revision, err = uc.apiClient.UpdatePreviewApplicationViews(ctx, appId, revision, newViews)
	if err != nil {
		return appId, revision, err
	}

	if err := uc.apiClient.DeployApplication(ctx, appId, revision); err != nil {
		return appId, revision, err
	}

	return appId, revision, nil
}
