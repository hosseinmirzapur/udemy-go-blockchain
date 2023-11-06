package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"

	"github.com/hosseinmirzapur/goblockchain/utils"
)

type Transaction struct {
	senderPrivKey        *ecdsa.PrivateKey
	senderPubKeey        *ecdsa.PublicKey
	senderBlockchAddr    string
	recipientBlockchAddr string
	value                float32
}

func NewTransaction(
	privKey *ecdsa.PrivateKey,
	pubKey *ecdsa.PublicKey,
	sender, recipient string,
	value float32,
) *Transaction {
	return &Transaction{privKey, pubKey, sender, recipient, value}
}

func (t *Transaction) GenSign() *utils.Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256(m)
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivKey, h[:])
	return &utils.Signature{r, s}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender"`
		Recipient string  `json:"recipient"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchAddr,
		Recipient: t.recipientBlockchAddr,
		Value:     t.value,
	})
}
