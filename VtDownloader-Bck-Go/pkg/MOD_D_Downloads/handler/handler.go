package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"VtDownloader.Back.Go/pkg/MOD_D_Downloads/messages"
	"VtDownloader.Back.Go/pkg/MOD_D_Downloads/models"
	"VtDownloader.Back.Go/pkg/MOD_D_Downloads/service"
	"VtDownloader.Back.Go/pkg/common"
	"github.com/gorilla/mux"
)

type Controller struct {
	service service.DownloadServiceI
}

func New(r *mux.Router, Service service.DownloadServiceI) *Controller {
	newController := new(Controller)
	newController.service = Service

	for _, route := range newController.Routes() {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	return newController
}

func (c *Controller) Routes() []common.Route {
	return []common.Route{
		{
			Path:    "/download/file",
			Handler: c.DownloadFile,
			Method:  http.MethodPost,
		},
	}
}

func (c *Controller) DownloadFile(w http.ResponseWriter, r *http.Request) {
	const Endpoint = "Download File"
	var downloadFileRequest models.DownloadFileRequest
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		println(Endpoint + "- Error reading body.")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.CommonErrorMessage{ErrorMessage: messages.ReadingBodyError})
	}
	err = json.Unmarshal(body, &downloadFileRequest)
	if err != nil {
		println(Endpoint + "- Error decoding body.")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.CommonErrorMessage{ErrorMessage: messages.DecodingBodyError})
	}
	if downloadFileRequest.DownloadId <= 0 {
		println(Endpoint + "- Error decoding body.")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.CommonErrorMessage{ErrorMessage: messages.InvalidDownloadId})
	} else if downloadFileRequest.DownloadURL == "" {
		println(Endpoint + "- Error decoding body.")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(common.CommonErrorMessage{ErrorMessage: messages.InvalidDownloadURL})
	}
	downloadFileResponse, err := c.service.StartDownloadFile(downloadFileRequest)
	if err != nil {
		println(Endpoint + "- Error downloading file " + downloadFileRequest.DownloadURL)
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(downloadFileResponse)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(downloadFileResponse)
}
