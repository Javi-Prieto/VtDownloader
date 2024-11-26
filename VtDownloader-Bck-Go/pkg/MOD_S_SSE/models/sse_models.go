package models

type FileDownloadedResponse struct {
	DownloadId string `json:download_id`
	Status     string `json:status`
}
