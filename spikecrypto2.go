package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	input := os.Args[1]

	output, err := signMessage(input)
	checkError(err)

	fmt.Println(output)

}

func signMessage(input string) (string, error) {
	_, err := os.Stat(fullPath("keypair.txt"))
	if err != nil {
		if os.IsNotExist(err) {
			// Create file and key pair and use it (function)
			return createSaveKey(input)
		} else {
			return "", err
		}
	} else {
		// Open file and get data in file and use it (function)
		return useKey(input)
	}
}

func fullPath(name string) string {
	var dir = fmt.Sprintf("%s/.local/share/signer", os.Getenv("HOME"))

	err := os.MkdirAll(dir, 0700)
	checkError(err)

	fileName := name
	fullPath := path.Join(dir, fileName)
	return fullPath
}

func shaSum(input string) []byte {
	inputShaSum := sha256.Sum256([]byte(input))

	inputSlice := make([]byte, len(inputShaSum))

	for i := range inputShaSum {
		inputSlice[i] = inputShaSum[i]
	}

	return inputSlice
}

func createSaveKey(input string) (string, error) {

	file, err := os.OpenFile(fullPath("keypair.txt"), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return "", err
	}
	defer file.Close()

	privateKey := new(ecdsa.PrivateKey)
	privateKey, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return "", err
	}

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, shaSum(input))
	if err != nil {
		return "", err
	}

	sign := r.Bytes()
	sign = append(sign, s.Bytes()...)

	encSign := base64.StdEncoding.EncodeToString(sign)

	var pubKey ecdsa.PublicKey
	pubKey = privateKey.PublicKey

	pemPubSlice, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		return "", err
	}

	var pemPubKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pemPubSlice}

	encPubPem := pem.EncodeToMemory(pemPubKey)

	pemPrivSlice, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return "", err
	}
	var pemPrivKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pemPrivSlice}

	encPrivPem := pem.EncodeToMemory(pemPrivKey)

	file.WriteString(string(encPrivPem))
	file.WriteString(string(encPubPem))

	var printOut output
	printOut.Message = input
	printOut.Signature = encSign
	printOut.PubKey = string(encPubPem)

	printJSON, err := json.MarshalIndent(printOut, "", "    ")
	if err != nil {
		return "", err
	}

	return string(printJSON), nil
}

func useKey(input string) (string, error) {
	file, err := os.Open(fullPath("keypair.txt"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	contents, err := ioutil.ReadFile(fullPath("keypair.txt"))
	if err != nil {
		return "", err
	}

	block, rest := pem.Decode(contents)

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, shaSum(input))
	if err != nil {
		return "", err
	}

	sign := r.Bytes()
	sign = append(sign, s.Bytes()...)

	encSign := base64.StdEncoding.EncodeToString(sign)

	var printOut output
	printOut.Message = input
	printOut.Signature = encSign
	printOut.PubKey = string(rest)

	printJSON, err := json.MarshalIndent(printOut, "", "    ")
	if err != nil {
		return "", err
	}

	return string(printJSON), nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type output struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	PubKey    string `json:"pubkey"`
}
