<p align="center">
  <img src="VtDwownloader-DkApp/VtDownloaderLogo.png" width="30%"/>
  <br>
</p>

# VtDownloader
[![GitHub release](https://img.shields.io/badge/Javier_Prieto-VtDownloader-blue?logo=github)](https://github.com/Javi-Prieto/VtDownloader/releases/latest)
[![GitHub alert](https://img.shields.io/badge/Alfa_Version-!-red)]()

This is a Desktop App used to download files from the internet but checking the url and the file downloaded in Virus Total. It uses [aria2](https://github.com/aria2/aria2) for the downloads.

## How to Download?
There is two options the first one is to copy the link of the download paste it on the Field and hit the button, that way you will start the download manually.

The other option is to use the firefox extension developed for this project and that you can find here, [extension](https://addons.mozilla.org/en-US/firefox/addon/vtdownloader_ext/).

>All the downloads are stored on your Downloads folder.

## Stack
<p align="center">
  <img src="https://skillicons.dev/icons?i=electron,go,python,javascript,html,css,figma" alt="Languages & Tools">
</p>

### Electron (Html, Css, Js)
I selected this framework for the frontend because it was faster to develop and obtain prettier results than other options.

### Go
I selected this language for my final version of the application because it was easier to package it in only one exe file, and also it supports much better threading than python

### Python
On a first time i select this language for my backend because of the simplicity of coding but i found out that was not what I need at all, despite of that it has a functional version but not a packaged one.

## Build
To build this application there is two options depending on the backend.
### Package Front
Run this
```shell
git clone https://github.com/Javi-Prieto/VtDownloader/
cd VtDownloader-DkApp
npm i
npm start
```
With that you will have the frontend running if you want to package it run `npm run package` and you will find a /dist folder with the installer

### Backend Go
First of all download aria2c from (https://github.com/aria2/aria2) and locate the aria2c.exe inside VtDownloader-Bck-Go directory. Next execute this commands
```shell
cd VtDownloader-Bck-Go/cmd
go run .
```

And you will have the go backend running. If you want to package it run `go build` instead and it will generate a cmd.exe file that you can run and use

### Backend Python
First of all download aria2c from (https://github.com/aria2/aria2) and locate the aria2c.exe inside VtDownloader-Bck directory. Next execute this commands
```shell
cd VtDownloader-Bck
pip install -r ./requirements.txt
python3 ./main.py
```

And you will have the python backend running. Sadly there is no an easy way to compile this backend.

## Extension
You can download the FIREFOX extension in this link [extension](https://addons.mozilla.org/en-US/firefox/addon/vtdownloader_ext/). This extension is on early access too so please notify me with any bug.

## Whats Next?
As you can see this is an early version of the application so there will take some time to create the final version 1. Also im open to receive new ideas from all of you but these are mine goals:
- Connect with the Virus Total API to check the files.
- Improve the firefox extension with, at least, a simple frontend.
- Create the extension for more navigators.
- Insert a database to store the downloads and some preferences.
- Add preferences to select other directories to save the downloads and when do you want to delete the storage.
- Fix the maximize button.
- Implement massive downloads.
- Implement an option to auto upload the downloaded file to a cloud storage.
- Maybe a YT video.
- Your ideas...