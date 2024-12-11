package models

type DownloadFileRequest struct {
	DownloadURL     string `json:"download_url"`
	DownloadId      int    `json:"download_id"`
	DownloadAnyways bool   `json:"download_anyways"`
}

type DownloadFileResponse struct {
	Status            string `json:"status"`
	StatusDescription string `json:"status_description"`
	DownloadId        int    `json:"download_id"`
	DownloadName      string `json:"download_name"`
}
