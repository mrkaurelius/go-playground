package mysigner

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/pem"
	"fmt"
	"testing"
)

func TestSignAndVerify(t *testing.T) {
	tk := TestKey()
	m := []byte("merhabayalandunya")
	digest := sha256.Sum256(m) // input must be hashed message
	sig, err := rsa.SignPKCS1v15(nil, tk, crypto.SHA256, digest[:])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("sig: %x\n", sig)
}

func TestMySignerSign(t *testing.T) {
	m := []byte("merhabayalandunya")
	ps := NewTest()
	sig, err := ps.Sign(nil, m, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("sig: %x\n", sig)
}

// TODO sign cert with p11 signer
func TestSignX509Cert(t *testing.T) {
	ps := NewTest()
	cert, err := NewBasicCert(*ps)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("cert bytes: %x\n", cert)
	encCert := pem.EncodeToMemory(&pem.Block{Bytes: cert, Type: "CERTIFICATE"})
	fmt.Println(string(encCert))
}
