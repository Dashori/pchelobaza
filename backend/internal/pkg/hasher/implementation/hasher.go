package hasherImplementation

import (
	"backend/internal/pkg/hasher"
	"golang.org/x/crypto/bcrypt"
)

type bcryptHasher struct {
}

func NewBcryptHasher() hasher.Hasher {
	return &bcryptHasher{}
}

func (b *bcryptHasher) GetHash(stringToHash string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(stringToHash), bcrypt.DefaultCost)
}

func (b *bcryptHasher) CheckUnhashedValue(hashedString, unhashedString string) bool {
	res := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(unhashedString))
	return res == nil
}
