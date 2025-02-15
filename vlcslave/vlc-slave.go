package main

import (
	"eneller/vlc-slave/internal/utils"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os/exec"
)

func main() {
	slog.Info("VLC-Slave Server listening on port " + utils.Port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info(r.UserAgent())
		focusWindowOrLaunch()
		receiveStreamFromClient(*r)
	})
	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		slog.Info(r.UserAgent())
		focusWindowOrLaunch()
		//TODO handle failure if key cant be found
		src := r.Header.Values(utils.HeadernameStream)
		if len(src) > 0 {
			openStream(src[0])
		} else {
			slog.Error(fmt.Sprintf("Received Empty Field for HTTP Header `%s`", utils.HeadernameStream))
		}
	})
	// start Server
	log.Fatal(http.ListenAndServe(utils.Port, nil))
}

func focusWindowOrLaunch() {
	// attempt to focus an existing VLC window
	//NOTE might also use xdotools here or equivalent wayland tools
	cmd := exec.Command("wmctrl", "-a", utils.VLCWindowName)
	if err := cmd.Run(); err != nil {
		//TODO launch VLC here
		slog.Error(fmt.Sprintf("Failed to run command: `%s` ", cmd.String()))
	}
}

func receiveStreamFromClient(request http.Request) {
	//TODO supplement with information provided by vlc master
	openStream(request.RemoteAddr)
}

/*
opens a network stream in vlc using a provided url using the following:
"add <uri> to playlist and start playback:  ?command=in_play&input=<uri>&option=<option>"
from file:///usr/share/vlc/lua/http/requests/README.txt
*/
func openStream(url string) {
	//TODO run against lua interface to open existing stream from URL

}
