package cgiserver

import (
	"http"
	"http/cgi"
	"os"
	"path/filepath"
)

type CgiHandler struct {
	http.Handler
	Root       string
	DefaultApp string
	UseLangMap bool
	LangMap    map[string]string
}

func CgiServer() *CgiHandler {
	path, _ := filepath.Abs(".")
	return &CgiHandler{nil, path, "", true, map[string]string{}}
}

func (h *CgiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var isCGI bool
	file := filepath.FromSlash(path)
	if len(file) > 0 && os.IsPathSeparator(file[len(file)-1]) {
		file = file[:len(file)-1]
	}
	ext := filepath.Ext(file)
	bin, isCGI := h.LangMap[ext]
	file = filepath.Join(h.Root, file)

	f, e := os.Stat(file)
	if e != nil || f.IsDirectory() {
		if len(h.DefaultApp) > 0 {
			file = h.DefaultApp
		}
		ext := filepath.Ext(file)
		bin, isCGI = h.LangMap[ext]
	}

	if isCGI {
		var cgih cgi.Handler
		if h.UseLangMap {
			cgih = cgi.Handler{
				Path: bin,
				Dir:  h.Root,
				Root: h.Root,
				Args: []string{file},
				Env:  []string{"SCRIPT_FILENAME=" + file},
			}
		} else {
			cgih = cgi.Handler{
				Path: file,
				Root: h.Root,
			}
		}
		cgih.ServeHTTP(w, r)
	} else {
		if (f != nil && f.IsDirectory()) || file == "" {
			tmp := filepath.Join(file, "index.html")
			f, e = os.Stat(tmp)
			if e == nil {
				file = tmp
			}
		}
		http.ServeFile(w, r, file)
	}
}
