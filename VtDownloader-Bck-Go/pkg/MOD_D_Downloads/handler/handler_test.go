package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"VtDownloader.Back.Go/pkg/MOD_D_Downloads/handler"
	"VtDownloader.Back.Go/pkg/MOD_D_Downloads/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	serviceMocked *serviceMock
	api           *mux.Router
)

type serviceMock struct{ mock.Mock }

func (s *serviceMock) StartDownloadFile(models.DownloadFileRequest) (*models.DownloadFileResponse, error) {
	return &models.DownloadFileResponse{Status: "success", StatusDescription: "", DownloadId: 1, DownloadName: "Name"}, nil
}

func TestMain(m *testing.M) {
	api = mux.NewRouter()
	serviceMocked = new(serviceMock)
	_ = handler.New(api, serviceMocked)

	_ = m.Run()
}

func TestDownloadFileSuccess(t *testing.T) {
	requestBody, _ := json.Marshal(models.DownloadFileRequest{
		DownloadURL:     "",
		DownloadId:      1,
		DownloadAnyways: false,
	})
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8000/download/file", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	serviceMocked.On("StartDownloadFile").Once()
	response := httptest.NewRecorder()
	api.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDownloadFileBadReq(t *testing.T) {
	requestBody, _ := json.Marshal(models.DownloadFileRequest{
		DownloadURL:     "",
		DownloadId:      -1,
		DownloadAnyways: false,
	})
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8000/download/file", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	serviceMocked.On("StartDownloadFile").Once()
	response := httptest.NewRecorder()
	api.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}
