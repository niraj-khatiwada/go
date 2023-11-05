package http

//import (
//	"fmt"
//	"net/http"
//)
//
//type Person struct {
//	name string
//}
//
//func main() {
//	fmt.Println("Server started at port 9000")
//	if err := http.ListenAndServe(":9000", http.HandlerFunc(handler)); err != nil {
//		fmt.Println("Error encountered during server start", err)
//	}
//}
//
//func handler(w http.ResponseWriter, r *http.Request) {
//	//fmt.Println(r.Method, r.URL, r.Host, r.UserAgent(), r.Header)
//
//	// URL.Path always returns decoded value
//	// URL.RawPath returns the exact same encoded value. For eg. Type /contact%2f. Here %2f is encoded value of "/"
//	fmt.Println(r.URL.Path, r.URL.RawPath, r.URL.Query())
//	switch r.URL.Path {
//	case "/":
//		{
//			person := Person{name: "Niraj"}
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(200)
//			if _, err := fmt.Fprint(w, person); err != nil {
//				return
//			}
//			break
//		}
//	case "/contact":
//		{
//			contactHandler(w, r)
//			break
//		}
//	default:
//		{
//			w.WriteHeader(404)
//			if _, err := fmt.Fprint(w, "<h1>Page not found</h1>"); err != nil {
//				return
//			}
//		}
//	}
//}
//
//func contactHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println(r.Cookies())
//	fmt.Println(r.Cookie("csrf"))
//	w.WriteHeader(404)
//	if _, err := fmt.Fprint(w, "<h1>Contact Us </h1>"); err != nil {
//		return
//	}
//}
//
//// If we pass direct handler to listenAndServe, we don't need  this function at all
////func routes() {
////	http.HandleFunc("/", rootHandler)
////	//http.HandleFunc("/contact", contactHandler)
////}
