package secrets

// type EncryptionStore interface {
// 	// EncryptionKey string
// 	// Backend EncryptionBackend

// 	Encrypt(value string) (string, error)
// 	Decrypt(encryptedValue string) (string, error)
// 	Rotate(newKey, encrytedValue string) (string, error)
// }

// var EncryptionKeys []string
// var EncryptionProvider string

// func init() {
// 	tokens := os.Getenv("SOURCEGRAPH_CRYPT")
// 	EncryptionKeys = strings.Split(tokens, ',')
// 	enc := os.Getenv("SOURCEGRAPH_TOKENSTORE")

// 	// for now, only the database
// 	if enc == "" {
// 		EncryptionProvider = EncryptionStore
// 	} else if enc == "" {

// 	} else {

// 	}
// }
