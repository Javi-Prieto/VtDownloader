class DownloadFileRequest{
    constructor(urlToSend, download_id, download_anyways){
        this.urlToSend = urlToSend;
        this.download_id = download_id;
        this.download_anyways = download_anyways;
    }
    toJson(){
        return JSON.stringify({
            download_url: this.urlToSend,
            download_id : this.download_id,
            download_anyways : this.download_anyways,
        });
    }
}

module.exports = DownloadFileRequest;