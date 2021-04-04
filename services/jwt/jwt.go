package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// Jwt wraps the signing key and the issuer
type Jwt struct {
	secretKey         string
	issuer            string
	accessExpiration  int64 // minutes
	refreshExpiration int64 // minutes
}

// JwtClaim adds user id as a claim to the token
type JwtClaim struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func New(secretKey, issuer string, accessExpiration, refreshExpiration int64) *Jwt {
	return &Jwt{secretKey, issuer, accessExpiration, refreshExpiration}
}

func (j *Jwt) GenerateTokenPair(userID uuid.UUID) (accessToken, refreshToken string, err error) {
	id := userID.String()
	accessToken, err = j.generateToken(id, j.accessExpiration)
	if err != nil {
		return
	}

	refreshToken, err = j.generateToken(id, j.refreshExpiration)
	return
}

// generates a jwt token
func (j *Jwt) generateToken(userID string, expiration int64) (signedToken string, err error) {
	claims := &JwtClaim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(expiration)).Unix(),
			Issuer:    j.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.secretKey))
	return
}

// ValidateToken validates the jwt token
func (j *Jwt) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.secretKey), nil
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
