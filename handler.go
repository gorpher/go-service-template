package service

import (
	"net/http"
	"os"

	"github.com/bnkamalesh/webgo/v6"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fs, err := os.OpenFile("./static/index.html", os.O_RDONLY, 0600)
	if err != nil {
		webgo.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	info, err := fs.Stat()
	if err != nil {
		webgo.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out := make([]byte, info.Size())
	_, err = fs.Read(out)
	if err != nil {
		webgo.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(out)
}
