package wallet

type Wallet struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Address    string `json:"address"`
}
