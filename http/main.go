package http

//import (
//	"fmt"
//	"net/http"
//)
//
//// Handler This is for root(/) handler
//type Handler struct{}
//
//func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	if _, err := fmt.Fprint(w, "<h1>Hello World</h1>"); err != nil {
//		return
//	}
//}
//
//func main() {
//	handler := Handler{}
//	fmt.Println("Server started at port 9000")
//	if err := http.ListenAndServe(":9000", handler); err != nil {
//		fmt.Println("Error encountered during server start")
//	}
//}
//
//// If you want route level handler
//func handler(w http.ResponseWriter, r *http.Request) {
//	if _, err := fmt.Fprint(w, "<h1>Hello World</h1>"); err != nil {
//		return
//	}
//}
//
//func main() {
//	http.HandleFunc("/login", handler)
//	fmt.Println("Server started at port 9000")
//	if err := http.ListenAndServe(":9000", nil); err != nil {
//		fmt.Println("Error encountered during server start")
//	}
//}
