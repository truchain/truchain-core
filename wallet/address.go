package wallet

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"golang.org/x/crypto/ripemd160"
)

func GenerateECDSAPair() (*ecdsa.PrivateKey, crypto.PublicKey, error) {
	pv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	pb := pv.Public()

	return pv, pb, nil
}

func EncodeECDSA(privateKey *ecdsa.PrivateKey, publicKey *crypto.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

func DecodeECDSA(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *crypto.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*crypto.PublicKey)

	return privateKey, publicKey
}

func GenerateSHA256Hash(data []byte) []byte {
	hash := sha256.Sum256(data)

	return hash[:]
}

func GenerateRipemd160Hash(data []byte) []byte {
	hasher := ripemd160.New()
	hasher.Write(data)
	hashBytes := hasher.Sum(nil)

	return hashBytes
}

func GenerateWallet() (*Wallet, error) {
	pv, pb, err := GenerateECDSAPair()
	if err != nil {
		return nil, err
	}
	pvs, pbs := EncodeECDSA(pv, &pb)

	// Generate Address
	address := GenerateSHA256Hash([]byte(pbs))
	address = GenerateSHA256Hash(address)
	address = GenerateRipemd160Hash(address)
	address = []byte(base64.StdEncoding.EncodeToString(address))

	response := new(Wallet)
	response.PrivateKey = pvs
	response.PublicKey = pbs
	response.Address = string(address)

	return response, nil
}
