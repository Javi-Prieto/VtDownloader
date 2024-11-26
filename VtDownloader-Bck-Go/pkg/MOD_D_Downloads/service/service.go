package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"VtDownloader.Back.Go/pkg/MOD_D_Downloads/messages"
	"VtDownloader.Back.Go/pkg/MOD_D_Downloads/models"
	modsse_service "VtDownloader.Back.Go/pkg/MOD_S_SSE"
	"github.com/gen2brain/beeep"
)

type DownloadServiceI interface {
	StartDownloadFile(downloadFileRequest models.DownloadFileRequest) (downloadFileResponse *models.DownloadFileResponse, err error)
}

type DownloadService struct{}

func New() DownloadServiceI {
	return new(DownloadService)
}

var (
	aria2cPath = "\\..\\aria2c.exe"
)

/*
Method used to start the download of the file

Params:

- downloadFileRequest, models.DownloadFileRequest: Store all the data required to make the download.

Return:

- downloadFileResponse, *models.DownloadFileResponse: Contains the data required to be return as a JSON response

- err, Error: Check if an error occurs during the process, if there is no error should be nil.
*/
func (m *DownloadService) StartDownloadFile(downloadFileRequest models.DownloadFileRequest) (downloadFileResponse *models.DownloadFileResponse, err error) {

	downloadName := strings.Split(downloadFileRequest.DownloadURL, "/")[len(strings.Split(downloadFileRequest.DownloadURL, "/"))-1]
	DownloadsPath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	filePath, err := os.Getwd()
	if err != nil {
		downloadFileResponse = &models.DownloadFileResponse{
			Status:            "failed",
			StatusDescription: messages.DownloadFileErrorMessage,
			DownloadId:        downloadFileRequest.DownloadId,
			DownloadName:      downloadName,
		}
		return nil, err
	}
	DownloadsPath += "/Downloads"
	if !strings.HasPrefix(downloadFileRequest.DownloadURL, "https://") {
		downloadFileResponse = &models.DownloadFileResponse{
			Status:            "failed",
			StatusDescription: messages.DownloadFileErrorMessage,
			DownloadId:        downloadFileRequest.DownloadId,
			DownloadName:      downloadName,
		}
		m.sendNotification(false, true, downloadFileRequest.DownloadURL)
		return downloadFileResponse, errors.New(messages.NotValidUrlFileError)
	}
	cmd := exec.Command(filePath+aria2cPath, "-d", DownloadsPath, downloadFileRequest.DownloadURL)
	downloadFileResponse = &models.DownloadFileResponse{
		Status:            "started",
		StatusDescription: messages.DownloadFileStartedMessage,
		DownloadId:        downloadFileRequest.DownloadId,
		DownloadName:      downloadName,
	}
	m.sendNotification(true, true, downloadFileRequest.DownloadURL)
	go m.downloadFile(cmd, downloadFileRequest.DownloadId, downloadFileRequest.DownloadURL)
	return downloadFileResponse, nil
}

/*
Method used to download the file, also it will send a notification with the result and response through sse server.

Params:

- command, *exec.Cmd: Is the command that will be executed.

- downloadId, int: Is the id of the download.

- downloadUrl, string: Is the string of the url that is going to be downloaded.
*/
func (m *DownloadService) downloadFile(command *exec.Cmd, downloadId int, downloadUrl string) {
	var message string
	err := command.Run()

	if err != nil {
		log.Fatalln("Error: " + err.Error())
		message = fmt.Sprintf("{\"download_id\": %d, \"status\": \"%s\"}", downloadId, "error")
		m.sendNotification(false, false, downloadUrl)
	} else {
		message = fmt.Sprintf("{\"download_id\": %d, \"status\": \"%s\"}", downloadId, "success")
		m.sendNotification(true, false, downloadUrl)
	}

	modsse_service.MessageChannel <- message
}

/*
Method used to notify the status of the download.

Params:

- isGood, bool: Check if the notification is good or if its a fail.

- isStart, bool: Check if the notification is the start of a download or a finish of one.
*/
func (m *DownloadService) sendNotification(isGood bool, isStart bool, downloadUrl string) {
	var (
		baseMessage = "Your download with URL: " + downloadUrl
	)
	filePath, _ := os.Getwd()
	println("FILE PATH ::: " + filePath)
	if !isGood {
		err := beeep.Notify("Your Download Has Failed", baseMessage+"has failed.", filePath+"\\..\\assets\\VtDownloaderLogo.ico")
		if err != nil {
			panic(err)
		}
	} else {
		if isStart {
			err := beeep.Notify("Your Download Has Started Correctly", baseMessage+" start.", filePath+"\\..\\assets\\VtDownloaderLogo.ico")
			if err != nil {
				panic(err)
			}
		} else {
			err := beeep.Notify("Your Download is complete", baseMessage+" has been completed.", filePath+"\\..\\assets\\VtDownloaderLogo.ico")
			if err != nil {
				panic(err)
			}
		}
	}

}
