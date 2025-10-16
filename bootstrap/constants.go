package bootstrap

type Constants struct {
	Context     Context
	JWTKeysPath JWTKeysPath
}

type Context struct {
	Translator string
	ID         string
}

type JWTKeysPath struct {
	PublicKey  string
	PrivateKey string
}

func NewConstants() *Constants {
	return &Constants{
		Context: Context{
			Translator: "translator",
			ID:         "id",
		},
		JWTKeysPath: JWTKeysPath{
			PublicKey:  "./internal/infrastructure/jwt/public_key.pem",
			PrivateKey: "./internal/infrastructure/jwt/private_key.pem",
		},
	}
}
