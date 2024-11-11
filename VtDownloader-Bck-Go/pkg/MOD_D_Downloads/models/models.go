package models

type DownloadFileRequest struct {
	DownloadURL     string `json:download_url`
	DownloadId      string `json:download_id`
	DownloadAnyways bool   `json:download_anyways`
}
