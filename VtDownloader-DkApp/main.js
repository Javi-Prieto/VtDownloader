const { app, BrowserWindow, ipcMain } = require('electron')
const path = require('node:path');

let window;
const createWindow = () => {
    const win = new BrowserWindow({
      minWidth: 1366 ,
      minHeight: 768,
      autoHideMenuBar: true,
      frame: false,
      center: true,
      backgroundColor: '#FFF',
      webPreferences: {
          nodeIntegration: true,
          preload: path.join(__dirname, '/src/preload.js'),
      },
      icon: './assets/VtDownloaderLogo.ico', 
    });
    win.loadFile('./src/index.html');
    return win;
}

app.whenReady().then(() => {
  window= createWindow();
});

ipcMain.on('close-app', function (){
  app.quit();
});

ipcMain.on('minimize-app', function (){
  window.minimize();
});

ipcMain.on('unmaximize-app', function (){
  window.unmaximize();
});

ipcMain.on('maximize-app', function (){
  window.maximize();
});