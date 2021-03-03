package access_token

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstant(t *testing.T) {
	// // use standard testing
	// if expirationTime != 24 {
	// 	t.Error(fmt.Sprintf("Expiration time constant must be 24. As current expirationTime constant is %v", expirationTime))
	// }

	// use testify
	assert.EqualValues(t, 24, expirationTime, fmt.Sprintf("Expiration time constant must be 24. As current expirationTime constant is %v", expirationTime))

}
func TestGetNewAccessToken(t *testing.T) {

	at := GetNewAccessToken()

	// // use standard testing
	// if at.isExpired() {
	// 	t.Error("access token should not be expired")
	// }

	// use testify
	assert.False(t, at.isExpired(), "access token should not be expired")

	// // use standard testing
	// if at.AccessToken != "" {
	// 	t.Error("new access token should not have define access token id")
	// }

	// use testify
	assert.EqualValues(t, "", at.AccessToken, "New access token sholud not defind access token id")

	// // use standard testing
	// if at.UserId != 0 {
	// 	t.Error("access token should not have an associated user id")
	// }

	// use testify
	assert.True(t, at.UserId == 0, "New access token should not associate user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	// // user standard testing
	// if !at.isExpired() {
	// 	t.Error("Empty access token should be expired by default")
	// }

	// user testify
	assert.True(t, at.isExpired(), "Empty access token should be expired by default")

	// // use standard testing
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	// if at.isExpired() {
	// 	t.Error("new access token is created and 3 hrs from now shold not be expired")
	// }

	// use testify
	assert.False(t, at.isExpired(), "new access token is created and 3 hrs from now shold not be expired")

}
