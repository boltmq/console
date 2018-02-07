// Copyright 2017 luoji

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const userAuthKey int = 0

var jwtSecret = []byte("secret")

type user struct {
	UserID   uint64 `json:"userID"`
	UserName string `json:"userName"`
	IsAdmin  bool   `json:"isAdmin"`
}

type userClaims struct {
	user
	jwt.StandardClaims
}

type authenticator struct {
}

func (auth *authenticator) Chain(w http.ResponseWriter, r *http.Request, ctx *Context) bool {
	// extract jwt
	authorizationHeader := r.Header.Get("Authorization")
	authRegex, _ := regexp.Compile("(?:Bearer *)([^ ]+)(?: *)")
	authRegexMatches := authRegex.FindStringSubmatch(authorizationHeader)
	if len(authRegexMatches) != 2 {
		// didn't match valid Authorization header pattern
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return false
	}
	jwtToken := authRegexMatches[1]

	// parse tokentoken
	token, err := jwt.ParseWithClaims(jwtToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return false
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok || !token.Valid {
		http.Error(w, "not authorized", http.StatusUnauthorized)
		return false
	}

	ctx.ctx = context.WithValue(r.Context(), userAuthKey, claims.user)
	return true
}

type loginHandler struct {
}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userParam := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := decoder.Decode(&userParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if userParam.Username != "admin" || userParam.Password != "admin" {
		http.Error(w, "invalid login", http.StatusUnauthorized)
		return
	}

	//generate token
	expire := time.Now().Add(time.Hour * 1).Unix()
	// Create the Claims
	claims := userClaims{
		user: user{
			UserID:   1,
			UserName: userParam.Username,
			IsAdmin:  true,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire,
			Issuer:    "login",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(jwtSecret)

	//output token
	tokenResponse := struct {
		Token string `json:"token"`
	}{signedToken}
	json.NewEncoder(w).Encode(tokenResponse)
}
