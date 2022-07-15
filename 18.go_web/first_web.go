package main

import "net/http"

// func helloWorld(resp http.ResponseWriter, req *http.Request) {
// 	resp.Write([]byte("hello world\n"))
// }

/* func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello world"))
} */
func helloWorld(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(resp, req)
		return
	}
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.Write([]byte(`{"hello", "world"}`))
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
