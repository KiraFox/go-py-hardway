package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	input := os.Args[1]

	// Turns the input into a slice of bytes, runs a SHA256 hash algorithm,
	// then returns the SHA256 checksum of that slice,
	// and returns it as an arry of [size]byte
	inputShaSum := sha256.Sum256([]byte(input))

	fmt.Println("Here is input: ", input)

	// The %x tells the fmt package to convert the array of bytes from the
	// Sum function into base16 output.
	//fmt.Printf("Here is SHA Sum: %x\n", inputShaSum)
	//fmt.Println("Without formatting...", inputShaSum)

	// Make a new slice of bytes with length equal to the length of inputShaSum
	inputSlice := make([]byte, len(inputShaSum))

	// At each index in inputShaSum, set the value at the same index in inputSlice
	// equal to the value at that index in inputShaSum
	// inputSlice becomes a slice of bytes with the same data as inputShaSum and
	// is still a cryptographic hash
	for i := range inputShaSum {
		inputSlice[i] = inputShaSum[i]
	}

	// Intialize variable privateKey as a new "Private Key" using the ECDSA package
	privateKey := new(ecdsa.PrivateKey)
	// Generate public and private key pair using the elliptic curve P521
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	checkError(err)

	// Intialize variable pubKey as a "Public Key" using ECDSA package
	var pubKey ecdsa.PublicKey
	// Set pubKey to the Public Key that corresponds to the Private Key in the key
	// pair generated earlier by privateKey
	pubKey = privateKey.PublicKey

	//fmt.Printf("This is private key: %x\n", privateKey)
	//fmt.Printf("This is public key: %x\n", pubKey)

	// The Sign function requires a io.Reader type(rand.Reader),
	// a *PrivateKey type(privateKey), and a hashed []byte type (inputSlice)
	// and returns 2 *big.Int types(r & s) and possibly an error (err)
	// The r & s variables together make the signature of the private key
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, inputSlice)
	checkError(err)

	// Initialize variable sign and set it to the conversion of r into a slice of bytes
	// Then add to sign the conversion of s into slices of bytes
	sign := r.Bytes()
	sign = append(sign, s.Bytes()...)

	//fmt.Printf("This is the signature: %x\n", sign)

	// Convert sign (slice of bytes) to a string using Base64 standard encoding
	// and set it to  variable encSign.  This gives you an encoded signature
	// of the signature you created before (r &s, then converted and set to sign)
	// It is easier to work with (print/save/etc) an encoded signature than the inital.
	encSign := base64.StdEncoding.EncodeToString(sign)

	fmt.Println("This is encoded signature of the input: ", encSign)

	// The MarshalPKIXPublicKey function from the x509 package requires a pointer
	// to a public key(&pubKey) then serialises it to DER-encoded PKIX format
	// which is returned as a slice of bytes(pemSlice) or returns an error (err)
	pemSlice, err := x509.MarshalPKIXPublicKey(&pubKey)
	checkError(err)

	// Create PEM encoded structure(Block) with the form:
	/*
		-----BEGIN Type-----
		base64-encoded Bytes
		-----END Type-----
	*/
	// where Type = "PUBLIC KEY" and the bytes to be encoded to base64 are pemSlice
	var pemKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pemSlice}

	// This sets the pemKey to a variable encPem to be used later
	encPem := pem.EncodeToMemory(pemKey)

	fmt.Println("EncPem: ", string(encPem))

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
