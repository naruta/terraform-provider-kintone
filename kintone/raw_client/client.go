package raw_client

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
)

type ApiClient struct {
	config ApiClientConfig
}

type ApiClientConfig struct {
	Host     string
	User     string
	Password string
}

func New(config ApiClientConfig) *ApiClient {
	return &ApiClient{config}
}

func (c *ApiClient) Call(ctx context.Context, req ApiRequest, resp interface{}) error {
	json, err := EncodeJson(req.Json)
	if err != nil {
		return err
	}

	url := req.Scheme + "://" + path.Join(c.config.Host, req.Path)
	httpRequest, err := http.NewRequest(req.Method, url, bytes.NewReader(json))
	if err != nil {
		return err
	}
	httpRequest = httpRequest.WithContext(ctx)
	httpRequest.Header.Set("X-Cybozu-Authorization", encodeBase64(strings.Join([]string{c.config.User, c.config.Password}, ":")))
	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		return errors.Errorf("status code: %s, body: %s", strconv.Itoa(httpResponse.StatusCode), string(body))
	}

	if err := DecodeJson(body, &resp); err != nil {
		return err
	}

	return nil
}

func encodeBase64(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

type ApiRequest struct {
	Scheme string
	Method string
	Path   string
	Json   interface{}
}
