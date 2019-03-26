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

	log.Println(path)

	switch path {
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
		serveReverseProxy("http://rcgechvalidator:8090", w, r)
	} else {
		http.Redirect(w, r, "https://www.ge.ch/dossier/geneve-numerique/blockchain", 308)
	}
}

func main() {
	keyPair, err := tls.LoadX509KeyPair("myservice.cert", "myservice.key")
	if err != nil {
		panic(err) // TODO handle error
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err) // TODO handle error
	}
	idpMetadataURL, err := url.Parse("http://ec2-18-184-234-216.eu-central-1.compute.amazonaws.com/ssorec.geneveid.ch_dgsi_blockchain.xml")
	//idpMetadataURL, err := url.Parse("http://ec2-18-184-234-216.eu-central-1.compute.amazonaws.com:8545/simplesaml/saml2/idp/metadata.php")
	if err != nil {
		panic(err) // TODO handle error
	}

	rootURL, err := url.Parse("http://ec2-18-184-234-216.eu-central-1.compute.amazonaws.com:8001")
	if err != nil {
		panic(err) // TODO handle error
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
		panic(err)
	}

	log.Println("HTTP running on 8080")
}
