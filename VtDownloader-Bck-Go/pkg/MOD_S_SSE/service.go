package modssse

import (
	"fmt"
	"net/http"
)

var (
	MessageChannel = make(chan string)
)

/*
Method used to send the notification to the front that the download is done

Params:

- isGood, bool: Check if the notification is good or if its a fail.

- isStart, bool: Check if the notification is the start of a download or a finish of one.
*/
func SseDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	defer func() {
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}()

	for {

		select {
		case msg := <-MessageChannel:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			fmt.Println("Client Disconnected")
			return
		}

	}

}
