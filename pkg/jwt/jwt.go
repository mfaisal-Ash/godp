package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	Exp    int64  `json:"exp"`
}

func secret() []byte {
	s := os.Getenv("JWT_SECRET")
	if s == "" {
		s = "dressmap-secret"
	}
	return []byte(s)
}

func b64(data []byte) string { return base64.RawURLEncoding.EncodeToString(data) }

func Generate(userID uint, role string) (string, error) {
	header := map[string]string{"alg": "HS256", "typ": "JWT"}
	claims := Claims{UserID: userID, Role: role, Exp: time.Now().Add(24 * time.Hour).Unix()}
	h, _ := json.Marshal(header)
	p, _ := json.Marshal(claims)
	unsigned := b64(h) + "." + b64(p)
	mac := hmac.New(sha256.New, secret())
	mac.Write([]byte(unsigned))
	return unsigned + "." + b64(mac.Sum(nil)), nil
}

func Parse(token string) (*Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token")
	}
	unsigned := parts[0] + "." + parts[1]
	mac := hmac.New(sha256.New, secret())
	mac.Write([]byte(unsigned))
	expected := b64(mac.Sum(nil))
	if !hmac.Equal([]byte(expected), []byte(parts[2])) {
		return nil, errors.New("invalid token signature")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	var claims Claims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, err
	}
	if time.Now().Unix() > claims.Exp {
		return nil, fmt.Errorf("token expired")
	}
	return &claims, nil
}
