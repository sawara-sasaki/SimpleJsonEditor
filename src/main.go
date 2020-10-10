package main

import (
	"io"
	"os"
	"fmt"
	"bytes"
	"os/exec"
	"runtime"
	"net/http"
	"encoding/json"
	"path/filepath"

	"github.com/sawara-sasaki/SimpleJsonEditor/src/action"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		os.Exit(1)
	}

	path := filepath.Dir(exe)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			if len(r.URL.Path[1:]) == 0 {
				http.ServeFile(w, r, path + "/static/index.html")
			} else {
				http.ServeFile(w, r, path + "/static/" + r.URL.Path[1:])
			}
		} else {
			body := r.Body
			defer body.Close()

			buf := new(bytes.Buffer)
			io.Copy(buf, body)
			var res action.ActionResponse
			var err error
			res, err = action.Handle(buf.Bytes());
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.Header().Set("Content-Type", "application/json")
				res.Status = http.StatusOK
				jsonBytes, _ := json.Marshal(res)
				w.Write(jsonBytes)
			}
		}
	})

	url := "http://localhost:8080"
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	http.ListenAndServe(":8080", nil)
}
