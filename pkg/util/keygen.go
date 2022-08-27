package util

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func GenerateRSAKey(keysize int) (publicKey string, privateKey string, error error) {
	key, err := rsa.GenerateKey(rand.Reader, keysize)
	if err != nil {
		return "", "", err
	}
	privateKey, err = savePEMKey(key)
	if err != nil {
		return "", "", err
	}

	publicKey, err = savePublicPEMKey(&key.PublicKey)
	if err != nil {
		return "", "", err
	}
	return publicKey, privateKey, nil
}

func savePEMKey(key *rsa.PrivateKey) (string, error) {
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	buf := bytes.NewBufferString("")
	err := pem.Encode(buf, privateKey)
	return buf.String(), err
}

func savePublicPEMKey(pubKey *rsa.PublicKey) (string, error) {
	pubByte, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", err
	}
	var pubkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubByte,
	}
	buff := bytes.NewBufferString("")
	err = pem.Encode(buff, pubkey)
	return buff.String(), err
}

func encodeECDSA(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

func GenerateECDSA256() (encPub string, encPriv string, err error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}
	publicKey := &privateKey.PublicKey

	encPriv, encPub = encodeECDSA(privateKey, publicKey)
	return encPub, encPriv, nil
}
