package hasher

type Hasher interface {
	GetHash(stringToHash string) ([]byte, error)
	CheckUnhashedValue(hashedString, unhashedString string) bool
}
