package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Jwt wraps the signing key and the issuer
type Jwt struct {
	accessSecret      string
	refreshSecret     string
	issuer            string
	accessExpiration  int64 // minutes
	refreshExpiration int64 // minutes
}

// JwtClaim adds user id as a claim to the token
type JwtClaim struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func New(issuer, accessSecret, refreshSecret string, accessExpiration, refreshExpiration int64) *Jwt {
	return &Jwt{issuer, accessSecret, refreshSecret, accessExpiration, refreshExpiration}
}

func (j *Jwt) GenerateTokenPair(userID uuid.UUID) (accessToken, refreshToken string, err error) {
	id := userID.String()
	accessToken, err = j.generateToken(id, j.accessSecret, j.accessExpiration)
	if err != nil {
		return
	}

	refreshToken, err = j.generateToken(id, j.refreshSecret, j.refreshExpiration)
	return
}

// generates a jwt token
func (j *Jwt) generateToken(userID, secret string, expiration int64) (signedToken string, err error) {
	claims := &JwtClaim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(expiration)).Unix(),
			Issuer:    j.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(secret))
	return
}

func (j *Jwt) ValidateAccessToken(accessToken string) (claims *JwtClaim, err error) {
	return validateToken(accessToken, j.accessSecret)
}

func (j *Jwt) ValidateRefreshToken(accessToken string) (claims *JwtClaim, err error) {
	return validateToken(accessToken, j.refreshSecret)
}

// ValidateToken validates the jwt token
func validateToken(signedToken, secret string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Failed to parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token has expired")
		return
	}

	return
}
