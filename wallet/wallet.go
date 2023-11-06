package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160" //this is deprecated but bitcoin uses it
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func NewWallet() *Wallet {
	w := &Wallet{}

	privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privKey
	w.publicKey = &privKey.PublicKey

	h2 := sha256.New()
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[:1], digest3[:])
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	chsum := digest6[:4]
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])
	address := base58.Encode(dc8)
	w.blockchainAddress = address

	return w
}

func (w *Wallet) PrivKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *Wallet) PrivKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PubKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *Wallet) PubKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}

func (w *Wallet) BlockchAddr() string {
	return w.blockchainAddress
}
