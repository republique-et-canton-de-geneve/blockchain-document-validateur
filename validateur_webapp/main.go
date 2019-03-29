package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"github.com/crewjam/saml/samlsp"
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
	path := r.URL.Path[1:]

	indexToServe := path

	log.Println("path", path)

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

	_, err := ioutil.ReadFile("mockup/"+string(indexToServe))

	if err == nil {
		http.ServeFile(w, r, "mockup/"+string(indexToServe))
	} else if strings.Split(path, "/")[0] == "api" {
		r.URL.Path = strings.TrimLeft(r.URL.Path, "api/")

		apiHost := os.Getenv("API_HOST")

		serveReverseProxy("http://"+apiHost, w, r)
	} else {
		http.Redirect(w, r, "https://www.ge.ch/dossier/geneve-numerique/blockchain", 308)
	}
}

func main() {
	keyName := os.Getenv("KEY_NAME")

	keyPair, err := tls.LoadX509KeyPair(keyName+".cert", keyName+".key")
	if err != nil {
		log.Fatal(err) // TODO handle error
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		log.Fatal(err) // TODO handle error
	}

	idpEnv := os.Getenv("IDP_METADATA")

	idpMetadataURL, err := url.Parse(idpEnv)
	if err != nil {
		log.Fatal(err) // TODO handle error
	}

	spEnv := os.Getenv("SP_URL")


	rootURL, err := url.Parse(spEnv)
	if err != nil {
		log.Fatal(err) // TODO handle error
	}

	samlSP, _ := samlsp.New(samlsp.Options{
		URL:            *rootURL,
		Key:            keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:    keyPair.Leaf,
		IDPMetadataURL: idpMetadataURL,
	})

	http.Handle("/saml/", samlSP)
	http.Handle("/", samlSP.RequireAccount(http.HandlerFunc(new(RouteHandler).ServeHTTP)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP running on 8080")
}
