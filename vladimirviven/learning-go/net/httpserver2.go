package main

import "fmt"

import "net/http"

func main() {
   mux := http.NewServeMux()
   // It is worth noting that HTTP package makes default ServeMux 
   // It is not necessary to use ServeMux. Use http.HandleFunc
   // to map a path to a handler function as commented (see end).

   hello := func(resp http.ResponseWriter, req *http.Request) {
               resp.Header().Add("Content-Type", "text/html")
               resp.WriteHeader(http.StatusOK)
               fmt.Fprint(resp, "Hello from Above!")
   }
   
   goodbye := func(resp http.ResponseWriter, req *http.Request) {
                 resp.Header().Add("Content-Type", "text/html")
                 resp.WriteHeader(http.StatusOK)
                 fmt.Fprint(resp, "Goodbye from Below")
   }
   mux.HandleFunc("/hello", hello)
   mux.HandleFunc("/goodbye", goodbye)
   // http.HandleFunc("/hello", hello)
   // http.HandleFunc("/goodbye", goodbye)
   
   http.ListenAndServe(":4040", mux)
   // When you are using default Mux the start function is as below
   // http.ListenAndServe(":4040", nil)
}
