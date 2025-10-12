package jwt

import (
	"crypto/rsa"
	"os"
	"sync"

	"github.com/golang-jwt/jwt/v5"
)

type JWTKeyManager struct {
	privateKey     *rsa.PrivateKey
	publicKey      *rsa.PublicKey
	mutex          sync.RWMutex
	isLoaded       bool
	privateKeyPath string
	publicKeyPath  string
}

func NewJWTKeyManager() *JWTKeyManager {
	return &JWTKeyManager{}
}

func (k *JWTKeyManager) LoadKeys() {
	k.mutex.Lock()
	defer k.mutex.Unlock()

	privKeyBytes, err := os.ReadFile(k.privateKeyPath)
	if err != nil {
		panic("failed to read private key: " + err.Error())
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privKeyBytes)
	if err != nil {
		panic("failed to parse private key: " + err.Error())
	}

	publicKeyBytes, err := os.ReadFile(k.publicKeyPath)
	if err != nil {
		panic("failed to read public key: " + err.Error())
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		panic("failed to parse public key: " + err.Error())
	}

	k.privateKey = privateKey
	k.publicKey = publicKey
	k.isLoaded = true
}

func (k *JWTKeyManager) GetPrivateKey() *rsa.PrivateKey {
	k.mutex.RLock()
	defer k.mutex.RUnlock()

	if !k.isLoaded {
		k.LoadKeys()
		return k.privateKey
	}
	return k.privateKey
}

func (k *JWTKeyManager) GetPublicKey() *rsa.PublicKey {
	k.mutex.RLock()
	defer k.mutex.RUnlock()

	if !k.isLoaded {
		k.LoadKeys()
		return k.publicKey
	}
	return k.publicKey
}
