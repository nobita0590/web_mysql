package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	"io/ioutil"
	"github.com/nobita0590/web_mysql/config"
	"sync"
)

var (
	publicFile = config.FilePath + "/key.pub"
	privateFile = config.FilePath + "/key.priv"
	rsaPublic	interface{}
	rsaPrivate	*rsa.PrivateKey
)

func Init()  {
	buildRSA()
}

func GenerateKey(claims jws.Claims) (b []byte, err error) {
	mux := sync.Mutex{}
	mux.Lock()
	defer mux.Unlock()
	jwtObject := jws.NewJWT(claims, crypto.SigningMethodRS512)
	b, err = jwtObject.Serialize(rsaPrivate)
	return
}

func ValidateKey(key []byte) (claims jwt.Claims,err error) {
	mux := sync.Mutex{}
	mux.Lock()
	defer mux.Unlock()

	jwtObject,err := jws.ParseJWT(key)
	if err != nil{
		return
	}
	err = jwtObject.Validate(rsaPublic, crypto.SigningMethodRS512)
	if err != nil{
		return
	}
	claims = jwtObject.Claims()
	return
}

func buildRSA()  {
	derBytes, err := ioutil.ReadFile(publicFile)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(derBytes)
	rsaPublic, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	der, err := ioutil.ReadFile(privateFile)
	if err != nil {
		panic(err)
	}
	block2, _ := pem.Decode(der)
	rsaPrivate, err = x509.ParsePKCS1PrivateKey(block2.Bytes)
	if err != nil {
		panic(err)
	}
}

func generateKeyFile() {
	f1, err := os.Create(privateFile)
	if err != nil {
		panic(err)
	}
	f2, err := os.Create(publicFile)
	if err != nil {
		panic(err)
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	//fmt.Println(privateKey)
	if err != nil {
		panic(err)
	}

	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyDer,
	}
	pem.Encode(f1, &privateKeyBlock)

	publicKey := privateKey.PublicKey
	publicKeyDer, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}

	publicKeyBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   publicKeyDer,
	}
	pem.Encode(f2, &publicKeyBlock)

	f1.Close()
	f2.Close()
}
