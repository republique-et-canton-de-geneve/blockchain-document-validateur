package internal

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

var rcgechUrl = "http://ge.ch/hrcintapp/companyReport.action?companyOfrcId13=CH-%s&lang=%s"

var ErrNoExtract = errors.New("rc-ge-ch-pdf: no extract found or unknown error")

func getPDF(ein string, langcode string) ([]byte, error) {
	url := fmt.Sprintf(rcgechUrl, ein, langcode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get: %v", err)
	}

	defer resp.Body.Close()
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return nil, ErrNoExtract
		case tt == html.SelfClosingTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "input"
			if !isAnchor {
				continue
			}

			// Extract the rcentId value, if there is one
			ok, rcentId := getRcentId(t)
			if !ok {
				continue
			}

			//Fetch PDF with rcentId
			pdfData, err := fetchRcentIdPDF(rcentId, langcode)
			if err != nil {
				return nil, fmt.Errorf("fetchRcentIdPDF: %v", err)
			}
			return pdfData, nil
		}
	}
	return nil, ErrNoExtract
}

var rcgechPDFUrl = "http://rc.ge.ch/hrcintreport/createReport?rcentId=%s&lang=%s&order=R&rad=Y"

func fetchRcentIdPDF(rcentId, langcode string) ([]byte, error) {
	url := fmt.Sprintf(rcgechPDFUrl, rcentId, langcode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get: %v", err)
	}

	defer resp.Body.Close()
	var b bytes.Buffer

	_, err = b.ReadFrom(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("b.ReadFrom")
	}

	return b.Bytes(), nil
}

// func getPdfData(rcentId, langcode string)
// Helper function to pull the rcentId attribute from an exceprt
func getRcentId(t html.Token) (ok bool, rcentId string) {
	for _, a := range t.Attr {
		if a.Key == "name" && a.Val == "rcentId" {
			ok = true
		}
		if a.Key == "value" {
			rcentId = a.Val
		}
	}

	// TODO factorize logic
	if rcentId == "" {
		ok = false
	}
	if !ok {
		rcentId = ""
	}
	return
}
