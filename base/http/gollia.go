package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func helloTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "http http")
	w.Write([]byte("hello"))
}

func haveOne() {
	startHttpServer()
}

func startHttpServer() {
	router := mux.NewRouter()

	// router.HandleFunc("/hello", helloTask)

	curDir, _ := GetCurPath()
	PthSep := string(os.PathSeparator)
	filePath := curDir + PthSep + "file"

	router.PathPrefix("/file/").Handler(http.StripPrefix("/file/", http.FileServer(http.Dir(filePath))))

	srv := &http.Server{
		Handler:      router,
		Addr:         ":7000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := http.ListenAndServe(srv.Addr, httpMiddleware(router))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetCurPath() (dir string, err error) {
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Sprintf("%s, %s", os.Args[0], err)
		return "", err
	}

	dir = filepath.Dir(path)
	fmt.Println(dir)
	return dir, nil
}

// 中间件
func httpMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Origin, X-Requested-With, Content-Type, Accept, common")

		h.ServeHTTP(w, r)
		if r.Method == "OPTIONS" {
			return
		}
	})
}
