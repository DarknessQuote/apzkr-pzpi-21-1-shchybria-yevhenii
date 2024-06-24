package infrastructure

import (
	"devquest-server/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	Auth struct {
		Issuer        string
		Audience      string
		Secret        string
		TokenExpiry   time.Duration
		RefreshExpiry time.Duration
		CookieDomain  string
		CookiePath    string
		CookieName    string
	}

	JWTUser struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		RoleTitle string    `json:"role"`
	}

	TokenPairs struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	Claims struct {
		jwt.RegisteredClaims
	}
)

func InitAuthSettings(conf *config.Config) *Auth {
	return &Auth{
		Issuer:        conf.Auth.Issuer,
		Secret:        conf.Auth.Secret,
		Audience:      conf.Auth.Audience,
		TokenExpiry:   time.Duration(conf.Auth.TokenExpiry) * time.Minute,
		RefreshExpiry: time.Duration(conf.Auth.RefreshExpiry) * time.Hour,
		CookieDomain:  conf.Auth.CookieDomain,
		CookiePath:    conf.Auth.CookiePath,
		CookieName:    conf.Auth.CookieName,
	}
}

func (a *Auth) GenerateTokenPairs(user *JWTUser) (TokenPairs, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"typ": "JWT",
		"name": user.Username,
		"role": user.RoleTitle,
		"sub": user.ID.String(),
		"aud": a.Audience,
		"iss": a.Issuer,
		"iat": time.Now().UTC().Unix(),
		"exp": time.Now().UTC().Add(a.TokenExpiry).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),
		"iat": time.Now().UTC().Unix(),
		"exp": time.Now().UTC().Add(a.RefreshExpiry).Unix(),
	})

	signedAccessToken, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	signedRefreshToken, err := refreshToken.SignedString([]byte(a.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	tokenPairs := TokenPairs{
		AccessToken: signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	return tokenPairs, nil
}

func (a *Auth) GetTokenFromHeader(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	w.Header().Add("Vary", "Authorization")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("no auth header")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("invalid auth header")
	}

	stringToken := headerParts[1]

	token, err := jwt.Parse(stringToken, func (token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Domain: a.CookieDomain,
		Path: a.CookiePath,
		Name: a.CookieName,
		Value: refreshToken,
		Expires: time.Now().Add(a.RefreshExpiry),
		MaxAge: int(a.RefreshExpiry.Seconds()),
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Secure: true,
	}
}

func (a *Auth) GetExpiredRefreshCookie() *http.Cookie {
	return &http.Cookie{
		Domain: a.CookieDomain,
		Path: a.CookiePath,
		Name: a.CookieName,
		Value: "",
		Expires: time.Unix(0, 0),
		MaxAge: -1,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Secure: true,
	}
}