const { contextBridge } = require('electron');
const homeBridge = require('./src/home/index_preload')

Bridge = homeBridge


contextBridge.exposeInMainWorld('Bridge', Bridge)


