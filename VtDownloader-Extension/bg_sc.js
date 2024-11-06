var downloadId = 0;

function saveDownloadIdAndCancelDonwload(downloadItem) {
  browser.downloads.cancel(downloadItem.id);
  downloadId++;
  sendDownloadToBackend(downloadItem.url);
}



browser.downloads.onCreated.addListener(saveDownloadIdAndCancelDonwload);

function sendDownloadToBackend(url) {
  fetch("http://localhost:8000/download/file", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      download_url: url,
      download_id: downloadId,
      download_anyways: false,
    }),
  })
    .then((response) => response.json())
    .then((data) => console.log("Download triggered:", data))
    .catch((error) => console.error("Error triggering download:", error));
}

