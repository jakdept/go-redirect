package main

import (
	"net/http"
	"flag"
)

listenPort := flag.Int("port", 80, "port to listen on")
tlsCert := flag.String("tlsCert", "", "If using TLS, cert to use")
tlsKey := flag.String("tlsKey", "", "If using TLS, key to use")
bindAddress := flag.String("bind", "0.0.0.0", "address to listen to")

resEncrypt := flag.Bool("https", true, "determine whether or not to include encryption in the redirect target")
resPort := flag.Int("targetPort", 443, "port to direct the request towards")
resCode := flag.Int("status", 301, "status code to return to the client")

func redir(w http.ResponseWriter, req *http.Request) {

	if resEncrypt {
		if resPort == 443 {
			http.Redirect(w, req, "https://"+req.Host+req.RequestURI, resCode)
		} else {
			http.Redirect(w, req, "https://"+req.Host:":"resPort+req.RequestURI, resCode)
		}
	} else {
		if resPort == 80 {
			http.Redirect(w, req, "http://"+req.Host+req.RequestURI, resCode)
		} else {
			http.Redirect(w, req, "http://"+req.Host:":"resPort+req.RequestURI, resCode)
		}
	}
}

func main() {
	flag.Parse()

	if tlsCert != "" && tlsKey != "" {
		if err := http.ListenAndServeTLS(string(bindAddress)+":"+listenPort, certFile, keyFile, handler); err != nil {
			log.Println(err)
		}
	} else {
		if err := http.ListenAndServe(string(bindAddress)+":"+listenPort, handler); err != nil {
			log.Println(err)
		}
	}
}
