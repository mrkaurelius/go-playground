package mysigner

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io"
)

type MySigner struct {
	key *rsa.PrivateKey
}

// omit errors
func New() *MySigner {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	return &MySigner{key: key}
}

func NewTest() *MySigner {
	return &MySigner{key: TestKey()}
}

// TODO constant key
func TestKey() (key *rsa.PrivateKey) {
	pemb, _ := pem.Decode([]byte(testKey))
	keyb := pemb.Bytes
	// fmt.Printf("keyb: %x\n", keyb)
	key, _ = x509.ParsePKCS1PrivateKey(keyb)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("key: %+v\n", key)
	return key
}

func (ps MySigner) Sign(rand io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	digest := sha256.Sum256(message)
	sig, err := rsa.SignPKCS1v15(rand, ps.key, crypto.SHA256, digest[:]) // input must be hashed message
	if err != nil {
		return nil, err
	}
	return sig, nil
}

// Although this type is an empty interface for backwards compatibility reasons, all public key types in the standard
// library implement the following interface
//
//	interface{
//	    Equal(x crypto.PublicKey) bool
//	}
func (ps MySigner) Public() *rsa.PublicKey {
	return &ps.key.PublicKey
	// return ps.key.Public difference?
}

// func (ps MySigner) PKR()
