package token

import (
	"crypto/ed25519"
	"encoding/hex"
)

var Pub ed25519.PublicKey
var Priv ed25519.PrivateKey

func init() {
	pubString, _ := hex.DecodeString("bb875c19e78919d43fcf163471a080ea7aff21c702e481006dd35e5eef70cb42")
	Pub = ed25519.PublicKey(pubString)

	privString, _ := hex.DecodeString("aaf303cb91974c2fe9aff0801cc69597680bddf014b00c61b1821a9c14f41cededa0c1410ed24edea135dd1bfa58203ae85e32a9c22f860b3c033b492c5fab43")
	Priv = ed25519.PrivateKey(privString)

}
