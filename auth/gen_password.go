package auth

import "golang.org/x/crypto/bcrypt"

func NormalizePassword(p string) []byte {
	return []byte(p)
}

func GeneratePassword(p string) string {
	nPass := NormalizePassword(p)

	hash, err := bcrypt.GenerateFromPassword(nPass, bcrypt.MinCost)
	if err != nil {
		return err.Error()
	}

	return string(hash)
}

func ComparePasswords(hashedPwd, inputPwd string) bool {
	nHash := NormalizePassword(hashedPwd)
	nInput := NormalizePassword(inputPwd)

	if err := bcrypt.CompareHashAndPassword(nHash, nInput); err != nil {
		return false
	}

	return true
}