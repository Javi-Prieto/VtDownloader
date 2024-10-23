const { app, BrowserWindow, ipcMain, net, shell } = require('electron')
const EventSource = require('eventsource');
const path = require('node:path');
const homeDir = require('os').homedir();
const DownloadFileRequest = require('./models/definitions');

var donwloadFileRequestClient;

let homeWindow;
var donwload_sse = new EventSource("http://localhost:8000/vt-sse/download");


const createWindow = () => {
    const win = new BrowserWindow({
      minWidth: 1366 ,
      minHeight: 768,
      autoHideMenuBar: false,
      frame: false,
      center: true,
      backgroundColor: '#FFF',
      webPreferences: {
          preload: path.join(__dirname, 'preload.js'),
          nodeIntegration: false,
          nodeIntegrationInWorker: true,
          contextIsolation: true,
      },
      icon: './assets/VtDownloaderLogo.ico', 
    });
    win.loadFile('./src/home/index.html');
    return win;
}

app.whenReady().then(() => {
  homeWindow= createWindow();
});

ipcMain.on('close-app', function (){
  ipcMain.removeAllListeners();
  app.quit();
});

ipcMain.on('minimize-app', function (){
  homeWindow.minimize();
});

ipcMain.on('unmaximize-app', function (){
  homeWindow.unmaximize();
});

ipcMain.on('maximize-app', function (){
  homeWindow.maximize();
});

donwload_sse.onmessage= function(event){
  console.log('Got', event.data);
  downloadFileBodyResponse = JSON.parse(event.data);
  homeWindow.webContents.send("downloaded-file-response", downloadFileBodyResponse);
};

ipcMain.handle("open-downloads", (event)=>{
  shell.openPath(path.join(homeDir, '/Downloads'));
});

ipcMain.handle("open-virustotal", (event)=>{
  shell.openExternal("https://www.virustotal.com/gui/home/upload");
});

ipcMain.handle("download-file",async (event, urlToSend, downloadId)=>{
  return new Promise((resolve,reject)=>{
    donwloadFileRequestClient = net.request({
      method: 'POST',
      protocol: 'http:',
      hostname: 'localhost',
      port: 8000,
      path:'/download/file',
      headers: ['Content-Type', 'application/json'],
    });
  
    let downloadBody = new DownloadFileRequest(urlToSend, downloadId, false);
    donwloadFileRequestClient.write(downloadBody.toJson(), 'utf-8');
    donwloadFileRequestClient.on('response', (response) => {
      var responseBody = '';
      response.on('data', (chunk) => {
          console.log(`BODY: ${chunk}`);
          responseBody += chunk.toString();
        });
      response.on('end', ()=>{
        try {
          let decodedBody = JSON.parse(responseBody); 
          console.log(`BODY DECODED: ${JSON.stringify(decodedBody)}`);
          resolve(decodedBody);
        } catch (error) {
          reject(new Error('Failed to parse response: ' + error.message));
        }
      });
    });
    donwloadFileRequestClient.on('error', (error) => {
        console.log(`ERROR: ${JSON.parse(error)}`);
        reject(new Error(`Request error: ${error}`));
    });
    donwloadFileRequestClient.end();
  });
});