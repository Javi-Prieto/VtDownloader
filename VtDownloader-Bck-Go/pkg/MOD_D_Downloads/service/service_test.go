package service_test

import (
	"testing"

	"VtDownloader.Back.Go/pkg/MOD_D_Downloads/models"
	"VtDownloader.Back.Go/pkg/MOD_D_Downloads/service"
	"github.com/stretchr/testify/assert"
)

var (
	serviceTest service.DownloadServiceI
	niceRequest = models.DownloadFileRequest{
		DownloadURL:     "https://go.dev/dl/go1.23.3.windows-amd64.msi",
		DownloadId:      1,
		DownloadAnyways: false,
	}
	badRequest = models.DownloadFileRequest{
		DownloadURL:     "http://go.dev/dl/go1.23.3.windows-amd64.msi",
		DownloadId:      2,
		DownloadAnyways: false,
	}
)

func TestMain(m *testing.M) {
	serviceTest = service.New()
}

func TestDownloadFileNice(t *testing.T) {
	response, err := serviceTest.StartDownloadFile(niceRequest)
	assert.NoError(t, err)
	assert.Equal(t, niceRequest.DownloadId, response.DownloadId)
}

func TestDownloadFileBad(t *testing.T) {
	response, err := serviceTest.StartDownloadFile(badRequest)
	assert.Error(t, err)
	assert.Equal(t, badRequest.DownloadId, response.DownloadId)
}
