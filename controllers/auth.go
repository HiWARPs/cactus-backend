package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/HiWARPs/cactus-backend/models"
	"github.com/HiWARPs/cactus-backend/utils"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

var jwtKey = []byte(os.Getenv("KEY"))

func Signup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		var user models.Admin

		json.NewDecoder(r.Body).Decode(&user)
		if len(user.Password) < 8 {
			http.Error(w, fmt.Sprintf("Password must have at least 8 characters"), 403)
			return
		}

		existingUser := &models.Admin{}
		err := mh.GetOneAdmin(existingUser, bson.M{"email": user.Email})
		if err == nil {
			http.Error(w, fmt.Sprintf("Something went wrong with the signup process"), 500)
			return
		}

		var errHash error
		user.Password, errHash = utils.GenerateHashPassword(user.Password)
		if errHash != nil {
			http.Error(w, fmt.Sprint(errHash), 500)
			return
		}

		_, err = mh.AddOneAdmin(&user)
		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		w.WriteHeader(201)
	}
}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Connect(); err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		user := &models.Admin{}

		json.NewDecoder(r.Body).Decode(&user)

		existingUser := &models.Admin{}
		err := mh.GetOneAdmin(existingUser, bson.M{"email": user.Email})
		if err != nil {
			http.Error(w, fmt.Sprintf("Not found"), 404)
			return
		}

		passwordErr := utils.CompareHashPassword(user.Password, existingUser.Password)
		if !passwordErr {
			http.Error(w, fmt.Sprintf("Invalid password"), 400)
			return
		}

		expirationTime := time.Now().Add(48 * time.Hour)

		claims := &models.Claims{
			StandardClaims: jwt.StandardClaims{
				Subject:   existingUser.Email,
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			http.Error(w, fmt.Sprint(err), 500)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			Path:     "/",
			Domain:   "localhost:8080",
			HttpOnly: false,
			SameSite: http.SameSiteLaxMode,
		})

		w.WriteHeader(200)
	}
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Now().Add(-50 * time.Hour),
			Path:     "/",
			Domain:   "localhost:8080",
			HttpOnly: false,
			SameSite: http.SameSiteLaxMode,
		})

		w.WriteHeader(200)
	}
}
