package main

import (
	"encoding/json"
	"github.com/gorilla/csrf"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type RouteHandler struct {
}

type TokenPayload struct {
	Token string `json:"token"`
}

/*
	Reverse Proxy Logic
*/

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

func (this *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if(r.Method != http.MethodGet && r.Method != http.MethodPost) {
		w.WriteHeader(444)
		return
	}

	mainURI := os.Getenv("MAIN_URI")

	path := r.URL.Path[1:]

	if strings.Split(path, "/")[0] != mainURI {
		http.Redirect(w, r, "https://www.ge.ch/dossier/geneve-numerique/blockchain", 308)
		return
	}

	path = strings.TrimPrefix(path, mainURI+"/")

	indexToServe := path

	// Switch to handle different languages
	switch path {
	case "":
		indexToServe = "index.fr.html"
	case "fr":
		indexToServe = "index.fr.html"
	case "en":
		indexToServe = "index.en.html"
	case "it":
		indexToServe = "index.it.html"
	case "de":
		indexToServe = "index.de.html"
	}

	_, err := ioutil.ReadFile("mockup/" + string(indexToServe))

	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	w.Header().Set("Content-Security-Policy", "default-src 'none'; script-src 'unsafe-inline' 'self'; connect-src 'self'; img-src data: *; style-src 'unsafe-inline' *; font-src *;")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	// If indexToServe is a valid file then return the file
	// Otherwise serve API if uri == /api/*
	// Finally redirect if incorrect request
	if err == nil {
		w.Header().Set("X-CSRF-Token", csrf.Token(r))

		http.ServeFile(w, r, "mockup/"+string(indexToServe))
	} else if strings.Split(path, "/")[0] == "api" {
		if (strings.Split(path, "/")[1] == "swagger.json") {
			w.WriteHeader(404)
		} else {
			w.Header().Set("X-CSRF-Token", csrf.Token(r))

			r.URL.Path = "/" + strings.TrimPrefix(r.URL.Path, "/"+mainURI+"/api/") // Remove api from uri

			apiHost := os.Getenv("API_HOST")

			serveReverseProxy("http://"+apiHost, w, r)
		}
	} else if strings.Split(path, "/")[0] == "token" {
		w.Header().Set("Content-Type", "application/json")

		token := csrf.Token(r)
		w.Header().Set("X-CSRF-Token", token)

		w.WriteHeader(200)

		payload := TokenPayload{Token: token}

		js, err := json.Marshal(payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)

	} else {
		http.Redirect(w, r, "https://www.ge.ch/dossier/geneve-numerique/blockchain", 308)
	}
}

func main() {
	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))

	// Main Gateway to Webapp & API, it needs SAML login
	http.Handle("/", http.HandlerFunc(CSRF(new(RouteHandler)).ServeHTTP))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP running on 8080")
}
