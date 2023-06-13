package mysigner

import (
	"crypto"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"math"
	"math/big"
	"time"
)

func NewBasicCert(signer MySigner) (cert []byte, err error) {
	serial, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))

	//The subject key identifier (SKID) is an x509 extension and thus actually part of the certificate.
	// The fingerprint instead is not part of the certificate but instead computed from the certificate.
	// A certificate does not need to have an SKID at all and can have at most one SKID.
	// security.stackexchange.com/q/200295
	skid, err := calculateSKID(signer.Public())
	if err != nil {
		return nil, err
	}

	template := &x509.Certificate{
		Subject: pkix.Name{
			CommonName: "Gumushane Yazilim Sanayi AS",
		},
		SerialNumber: serial,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(100, 0, 0),

		SubjectKeyId:          skid,
		AuthorityKeyId:        skid,
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLenZero:        true,
	}

	der, err := x509.CreateCertificate(rand.Reader, template, template, signer.Public(), signer.key)
	if err != nil {
		return nil, err
	}
	return der, nil
}

func calculateSKID(pubKey crypto.PublicKey) ([]byte, error) {
	spkiASN1, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	var spki struct {
		Algorithm        pkix.AlgorithmIdentifier
		SubjectPublicKey asn1.BitString
	}
	_, err = asn1.Unmarshal(spkiASN1, &spki)
	if err != nil {
		return nil, err
	}
	skid := sha1.Sum(spki.SubjectPublicKey.Bytes)
	return skid[:], nil
}
