const ipcRenderer = require('electron').ipcRenderer;
let isMax = false;

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
    
}

window.addEventListener('DOMContentLoaded', function() {
    wireUpButtons();
});
