package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

type Claims struct {
	UserID    uuid.UUID `json:"user_id"`
	SessionID uuid.UUID `json:"session_id"`
	Type      TokenType `json:"type"`
	jwt.RegisteredClaims
}

type JWTService struct {
	accessSecret  []byte
	refreshSecret []byte
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

func NewJWTService(accessSecret, refreshSecret string, accessTTL, refreshTTL time.Duration) *JWTService {
	return &JWTService{
		accessSecret:  []byte(accessSecret),
		refreshSecret: []byte(refreshSecret),
		accessTTL:     accessTTL,
		refreshTTL:    refreshTTL,
	}
}

func (s *JWTService) GenerateAccessToken(userID, sessionID uuid.UUID) (string, time.Time, error) {
	exp := time.Now().Add(s.accessTTL)
	claims := Claims{
		UserID:    userID,
		SessionID: sessionID,
		Type:      TokenTypeAccess,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.accessSecret)
	return signed, exp, err
}

func (s *JWTService) GenerateRefreshToken(userID, sessionID uuid.UUID) (string, time.Time, error) {
	exp := time.Now().Add(s.refreshTTL)
	claims := Claims{
		UserID:    userID,
		SessionID: sessionID,
		Type:      TokenTypeRefresh,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.refreshSecret)
	return signed, exp, err
}

func (s *JWTService) ParseAccessToken(tokenStr string) (*Claims, error) {
	return s.parseToken(tokenStr, s.accessSecret, TokenTypeAccess)
}

func (s *JWTService) ParseRefreshToken(tokenStr string) (*Claims, error) {
	return s.parseToken(tokenStr, s.refreshSecret, TokenTypeRefresh)
}

func (s *JWTService) parseToken(tokenStr string, secret []byte, expected TokenType) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}
	if claims.Type != expected {
		return nil, fmt.Errorf("invalid token type")
	}
	return claims, nil
}

func (s *JWTService) AccessTTL() time.Duration  { return s.accessTTL }
func (s *JWTService) RefreshTTL() time.Duration { return s.refreshTTL }

func GenerateOpaqueToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
