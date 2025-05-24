package auth

import "golang.org/x/crypto/bcrypt"

type BcryptService interface {
	Hash(pwd string) (string, error)
	Compare(pwd, hash string) bool
}

type bcryptSvc struct{}

func (s *bcryptSvc) Hash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(bytes), err
}

func (s *bcryptSvc) Compare(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

var bcryptSvcRef *bcryptSvc

func NewBcryptService() BcryptService {
	if bcryptSvcRef == nil {
		bcryptSvcRef = &bcryptSvc{}
	}
	return bcryptSvcRef
}
