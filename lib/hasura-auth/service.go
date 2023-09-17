package hasuraauth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5/middleware"
)

type Service struct {
	httpClient  *http.Client
	baseUrl     url.URL
	adminSecret string
}

func NewService(httpClient *http.Client, uri url.URL, a string) *Service {
	return &Service{
		httpClient:  httpClient,
		baseUrl:     uri,
		adminSecret: a,
	}
}

func (s *Service) Login(ctx context.Context, req RequestEmailPassword) (result ResponseEmailPassword, status int, err error) {
	reqSerialized, err := json.Marshal(req)
	if err != nil {
		return result, http.StatusBadRequest, errors.Join(errors.New("serialize request email password"), err)
	}
	path := s.baseUrl.JoinPath("/signin/email-password").String()
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewBuffer(reqSerialized))
	if err != nil {
		return result, http.StatusBadRequest, errors.Join(errors.New("create request"), err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Hasura-Admin-Secret", s.adminSecret)
	requestID := ctx.Value(middleware.RequestIDKey).(string)
	if requestID != "" {
		request.Header.Set(middleware.RequestIDHeader, requestID)
	}
	resp, err := s.httpClient.Do(request)
	if err != nil {
		return result, http.StatusServiceUnavailable, errors.Join(errors.New("send request"), err)
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return result, resp.StatusCode, errors.New("invalid status code")
	} else if resp.StatusCode >= 500 {
		return result, http.StatusServiceUnavailable, errors.New("server error")
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, http.StatusInternalServerError, errors.Join(errors.New("decode response"), err)
	}
	return result, resp.StatusCode, nil
}
