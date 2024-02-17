package hasher

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(passwd string) (string, error) {
	hPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hPasswd), nil
}

func CompareWithPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
