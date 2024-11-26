package main

import (
	"net/http"

	downloadsHandler "VtDownloader.Back.Go/pkg/MOD_D_Downloads"
	modssseService "VtDownloader.Back.Go/pkg/MOD_S_SSE"
	"github.com/gorilla/mux"
)

func main() {
	println("--- INITIATED THE VTDOWNLOADER BACK ---")
	router := mux.NewRouter()
	downloadsHandler.Init(router)
	router.HandleFunc("/vt-sse/download", modssseService.SseDownload).Methods(http.MethodGet)
	http.ListenAndServe(":8000", router)
}
