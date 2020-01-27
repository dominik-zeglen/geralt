package api

import (
	"context"
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dominik-zeglen/geralt/models"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthRequest struct {
	Email string `json:"email"`
}
type AuthResponse struct {
	Token string `json:"token"`
}

func (api *API) handleAuth(w http.ResponseWriter, r *http.Request) {
	var data AuthRequest
	var user models.User

	reqDecodeErr := json.NewDecoder(r.Body).Decode(&data)
	if reqDecodeErr != nil {
		http.Error(w, reqDecodeErr.Error(), http.StatusBadRequest)
		return
	}

	collection := api.db.Collection(models.UsersCollectionKey)
	err := collection.FindOne(context.TODO(), bson.M{
		"email": data.Email,
	}).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	claims := UserClaims{
		ID: user.ID.Hex(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(api.conf.secret))

	reply := AuthResponse{
		Token: tokenString,
	}

	jsonResponse, resEncodeErr := json.Marshal(&reply)
	if resEncodeErr != nil {
		http.Error(w, resEncodeErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
