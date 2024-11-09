const { contextBridge, ipcRenderer } = require('electron')

let isMax = false;

let homeBridge = {
    downloadFile: (urlToDownload, downloadId) => {
        return new Promise((resolve, reject) => {
            ipcRenderer.invoke("download-file", urlToDownload, downloadId)
                .then((body) => {
                    console.log(`DATA 1 ::::: ${body}`)
                    resolve(body);})
                .catch(err => reject(err));
        });
    },
    onDownloadFileResponse: (callback) =>{
        ipcRenderer.on("downloaded-file-response", (event, message)=>{
            callback(message);
        })
    },
    openFileExplorer: ()=>{
        ipcRenderer.invoke("open-downloads");
    },
    openVirusTotal: ()=>{
        ipcRenderer.invoke("open-virustotal");
    }
};

const wireUpButtons = () => {
    let closeBtn = document.getElementById('close-b');
    let minBtn = document.getElementById('min-b');
    let resBtn = document.getElementById('resize-b');
    closeBtn.addEventListener('click', function () {
        ipcRenderer.send('close-app');
    });
    
    minBtn.addEventListener('click', function(){
        ipcRenderer.send('minimize-app');
    });

    if(isMax){
        isMax = false;
        resBtn.addEventListener('click', function(){
            ipcRenderer.send('unmaximize-app');
        });
    }else{
        isMax = true;
        resBtn.addEventListener('click', function(){
            ipcRenderer.send('maximize-app');
        });
    }
    
};

window.addEventListener('DOMContentLoaded', function() {
    wireUpButtons();
});

contextBridge.exposeInMainWorld("homeBridge", homeBridge);
