package main

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// getCerts reads a certificate chain from a file and returns it as
// an array of certificates
func getCerts(filename string) ([]*x509.Certificate, error) {

	var certs []*x509.Certificate

	if filename == "" {
		return nil, errors.New("filename not specified")
	}

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	for len(data) > 0 {

		var block *pem.Block

		block, data = pem.Decode(data)

		if block == nil {
			return nil, errors.New("no PEM data is found")
		}

		if block.Type != "CERTIFICATE" {
			return nil, errors.New("not a certificate")
		}

		cert, err := x509.ParseCertificate(block.Bytes)

		if err != nil {
			return nil, err
		}

		certs = append(certs, cert)
	}

	return certs, nil
}

func main() {

	certs, err := getCerts("test.pem")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var clientCert *x509.Certificate
	var caCert *x509.Certificate

	// First cert has to be the cert of the client
	if len(certs) > 0 {
		if !certs[0].IsCA {
			clientCert = certs[0]
		}
	}

	// Second cert has to be the cert of the CA
	if len(certs) > 1 {
		if certs[1].IsCA {
			caCert = certs[1]
		}
	}

	// Print some data
	fmt.Printf("Version: %d\nCA is CA: %t\n", clientCert.Version, caCert.IsCA)
}
