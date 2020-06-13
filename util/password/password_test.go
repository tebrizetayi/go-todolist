package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	password := "123456"
	passwordHash, _ := HashPassword(password)

	//Test valid password
	assert.Nil(t, VerifyPassword(string(passwordHash), password))

	//Test invalid password
	FalsePassword := "1234567"
	assert.NotNil(t, VerifyPassword(string(passwordHash), FalsePassword))
}
