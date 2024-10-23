
let downloadId = 1;
jQuery(function() {
    $('#donwload_button').on('click', ()=>{
        console.log('CLICKED ::::');
        let urlToSend = document.getElementById('donwload_url').value;
        window.homeBridge.downloadFile(urlToSend, downloadId)
            .then((data)=>{
                console.log(`DATA:::${data}`)
                toShowFile = returnDownloadFileComponent('d-in-progress', data.download_name, data.download_id);
                $("#downloaded-files").append(toShowFile);
            }).catch((error)=>{
                console.log(`DATA:::${error}`)
                toShowFile =  returnDownloadFileComponent('d-failed', error.download_name, error.download_id);
                $("#downloaded-files").append(toShowFile);
            });
        $("#donwload_url").val("");
        downloadId ++;
    });

    $('#downloaded-files').on('click', '.search-file-button',()=>{
        window.homeBridge.openFileExplorer();
    });

    $('#downloaded-files').on('click', '.security-file-button',()=>{
        window.homeBridge.openVirusTotal();
    });

    $('#downloaded-files').on('click', '.delete-button',(event)=>{
        $(event.target).parent().parent().parent().remove();
    });

    window.homeBridge.onDownloadFileResponse((fileResponse)=>{
        console.log('DOWNLOAD ID :: ' + fileResponse.download_id);
        console.log('DOWNLOAD STATUS :: ' + fileResponse.status);
        downloadName = $(`#download-${fileResponse.download_id}-name`).text();
        $(`#download-${fileResponse.download_id}`).children().remove();
        $(`#download-${fileResponse.download_id}`).append(returnFileDownloadedChildrens(fileResponse.status, downloadName, downloadId));
    });

});


function returnFileDownloadedChildrens(status, downloadName, downloadId){
    if (status == 'success'){
        return `
            <img src="../../assets/downloaded-file.png" alt="Downloaded File Icon" id="download-${downloadId}-img">
            <h4>${downloadName}</h4>
            <div class="progress-loader">
                <div class="progress"></div>
            </div>
            <div id="d-file-buttons-container">
                <button type="button" class="security-file-button d-file-buttons">
                    <img src="../../assets/scan-file.png" alt="Scan Icon">
                </button>
                <button type="button" class="search-file-button d-file-buttons">
                    <img src="../../assets/directory-icon.png" alt="Search Icon" class='search-file-button-img'>
                </button>
                <button type="button" class="delete-button d-file-buttons" >
                    <img src="../../assets/delete-icon.png" alt="Delete Icon">
                </button>
            </div>
        `;
    }else{
        return `
            <img src="../../assets/failed-file.png" alt="Failed File Icon" id="download-${downloadId}-img">
            <h4 id="download-${downloadId}-name">${downloadName}</h4>
            <div class="progress-loader">
                <div class="progress"></div>
            </div>
            <div id="d-file-buttons-container">
                <button type="button" class="retry-icon-button d-file-buttons">
                    <img src="../../assets/retry_icon_disabled.png" alt="Scan Icon">
                </button>
                <button type="button" class="search-file-button d-file-buttons">
                    <img src="../../assets/directory-icon.png" alt="Search Icon">
                </button>
                <button type="button" class="delete-button d-file-buttons">
                    <img src="../../assets/delete-icon.png" alt="Delete Icon">
                </button>
            </div>
        `;
    }
    
}

function returnDownloadFileComponent(downloadStatus, downloadName, downloadId){
    switch (downloadStatus){
        case 'd-in-progress':
           return `
                <div class="file-d" id="download-${downloadId}">
                    <img src="../../assets/downloading-file.png" alt="" id="download-${downloadId}-img">
                    <h4 id="download-${downloadId}-name">${downloadName}</h4>
                    <div class="progress-loader">
                        <div class="progress"></div>
                    </div>
                    <div id="d-file-buttons-container">
                        <button type="button" class="security-file-button d-file-buttons">
                            <img src="../../assets/cancel-icon.png" alt="Scan Icon">
                        </button>
                        <button type="button" class="search-file-button d-file-buttons">
                            <img src="../../assets/directory-icon.png" alt="Search Icon">
                        </button>
                        <button type="button" class="delete-button d-file-buttons">
                            <img src="../../assets/delete-icon.png" alt="Delete Icon">
                        </button>
                    </div>
                </div>
            `;
        case 'd-failed':
            return `
                <div class="file-d" id="download-${downloadId}">
                    <img src="../../assets/failed-file.png" alt="Failed File Icon" id="download-${downloadId}-img">
                    <h4 id="download-${downloadId}-name">${downloadName}</h4>
                    <div class="progress-loader">
                        <div class="progress"></div>
                    </div>
                    <div id="d-file-buttons-container">
                        <button type="button" class="retry-icon-button d-file-buttons">
                            <img src="../../assets/retry_icon_disabled.png" alt="Scan Icon">
                        </button>
                        <button type="button" class="search-file-button d-file-buttons">
                            <img src="../../assets/directory-icon.png" alt="Search Icon">
                        </button>
                        <button type="button" class="delete-button d-file-buttons">
                            <img src="../../assets/delete-icon.png" alt="Delete Icon">
                        </button>
                    </div>
                </div>
            `;
        default:
            console.error("Not valid status.")
            return '';
    }
}
