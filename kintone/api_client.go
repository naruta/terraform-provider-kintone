package kintone

import (
	"context"
	"fmt"
	"strconv"
)

type ApiError struct {
	Method       string
	RequestPath  string
	StatusCode   int
	ResponseBody string
}

func (e ApiError) Error() string {
	return fmt.Sprintf("method: %s, path: %s, status code: %s, body: %s", e.Method, e.RequestPath, strconv.Itoa(e.StatusCode), e.ResponseBody)
}

type ApiClientConfig struct {
	Host     string
	User     string
	Password string
}

type ApiClient interface {
	CreatePreviewApplication(ctx context.Context, name string) (AppId, Revision, error)
	UpdatePreviewApplicationSettings(ctx context.Context, appId AppId, revision Revision, setting Setting) (Revision, error)
	UpdatePreviewApplicationStatus(ctx context.Context, appId AppId, revision Revision, status Status) (Revision, error)
	CreatePreviewApplicationFormFields(ctx context.Context, appId AppId, revision Revision, fields []Field) (Revision, error)
	UpdatePreviewApplicationFormFields(ctx context.Context, appId AppId, revision Revision, fields []Field) (Revision, error)
	DeletePreviewApplicationFormFields(ctx context.Context, appId AppId, revision Revision, fieldCodes []FieldCode) (Revision, error)
	UpdatePreviewApplicationViews(ctx context.Context, appId AppId, revision Revision, views []View) (Revision, error)
	FetchPreviewApplicationViews(ctx context.Context, appId AppId) ([]View, error)

	DeployApplication(ctx context.Context, appId AppId, revision Revision) error
	FetchApplication(ctx context.Context, appId AppId) (Application, error)
	DeleteApplication(ctx context.Context, appId AppId) error

	CreateRecord(ctx context.Context, appId AppId, record Record) (RecordId, RecordRevision, error)
	UpdateRecord(ctx context.Context, appId AppId, record Record) (RecordRevision, error)
	DeleteRecords(ctx context.Context, appId AppId, ids []RecordId) error
}
