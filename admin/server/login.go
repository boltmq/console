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
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

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
