package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/amosehiguese/restaurant-api/auth"
	"github.com/amosehiguese/restaurant-api/models"
	"github.com/amosehiguese/restaurant-api/store"
	"github.com/amosehiguese/restaurant-api/types"
)



func SignUp(w http.ResponseWriter, r *http.Request) {
	var input types.SignUp

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnprocessableEntity,
			"msg": "Unprocessable entity",
		})
		return
	}	

	validate := types.NewValidator()
	if err := validate.Struct(input); err != nil {
			json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": types.ValidatorErrors(err),
		})
		return
	}

	q := store.GetQuery()
	role, err := q.RetrieveRole(ctx, 2)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}

	
	user := models.CreateUserParams {
		FirstName: input.FirstName,
		LastName: input.LastName,
		Username: input.UserName,
		Email: input.Email,
		PasswordHash: auth.GeneratePassword(input.Password),
		UserRole: role.ID,
		CreatedAt: time.Now(),
	}

	if err := validate.Struct(user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": types.ValidatorErrors(err),
		})
		return
	}

	var newUser models.User
	if newUser, err = q.CreateUser(ctx, user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal Server Error",
		})
		return
	}

	newUser.PasswordHash = ""

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success":  true,
		"user": newUser,
	})

}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var input types.SignIn
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnprocessableEntity,
			"msg": "Unprocessable entity",
		})
		return
	}

	q := store.GetQuery()
	foundUser, err := q.RetrieveUserByEmail(ctx, input.Email)
	fmt.Println(foundUser.PasswordHash, err)
	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success":false,
			"msg": "user with the given email is not found",
		})
		return 
	}

	compareUserPassword := auth.ComparePasswords(foundUser.PasswordHash, input.Password)

	if !compareUserPassword {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"msg": "wrong user email address or password",
		})
		return
	}

	tokens, err := auth.GenNewTokens(foundUser.ID, foundUser.UserRole)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success":false,
			"msg": err.Error(),
		})
		return
	}

	userId := foundUser.ID
	

	connRedis, err := store.RedisConn()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success":false,
			"msg": err.Error(),
		})
		return 
	}

	errSaveToRedis := connRedis.Set(context.Background(),userId.String(), tokens.Refresh, 0).Err()
	if errSaveToRedis != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"msg": errSaveToRedis.Error(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true, 
		"tokens": resp{
			"access": tokens.Access,
			"refresh": tokens.Refresh,
		},
	})

}

func SignOut(w http.ResponseWriter, r *http.Request) {
	claims, err := auth.ExtractTokenMetadata(r)
	if err != nil {
		l.Error(err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal Server Error",
		})
		return
	}
	
	userID := claims.UserID.String()

	connRedis, err := store.RedisConn()
	if err != nil {
		l.Error(err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal Server Error",
		})
		return
	}
	
	errDelFromRedis := connRedis.Del(ctx, userID).Err()
	if errDelFromRedis != nil {
		l.Error(err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": errDelFromRedis.Error(),
		})
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"code": http.StatusNoContent,
	})
}

func RenewTokens(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Unix()

	claims, err := auth.ExtractTokenMetadata(r)
	if err != nil {
		l.Error(err.Error())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal Server Error",
		})
		return
	}

	expiresAccessToken := claims.EAT
	if now > expiresAccessToken {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnauthorized,
			"msg": "Unauthorized, check expiration time of your token",
		})
		return
	}

	var renew types.Renew
	err = json.NewDecoder(r.Body).Decode(&renew)
	if err != nil {
		l.Errorln(err)
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnprocessableEntity,
			"msg": "Unprocessable entity",
		})
		return
	}

	expiresRefreshToken, err := auth.ParseRefreshToken(renew.RefreshToken)
	if err != nil {
		l.Errorln(err)
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}

	if now < expiresRefreshToken {
		userID := claims.UserID

		q := store.GetQuery()
		user, err := q.RetrieveUser(ctx, userID)
		if err != nil {
			l.Errorln(err)
			json.NewEncoder(w).Encode(resp{
				"success": false,
				"code": http.StatusNotFound,
				"msg": fmt.Sprintf("User with the given ID %v is not found", userID),
			})
			return
		}

		tokens, err := auth.GenNewTokens(userID, user.UserRole)
		if err != nil {
			l.Error(err.Error())
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp{
				"success": false,
				"code": http.StatusInternalServerError,
				"msg": "Internal Server Error",
			})
			return
		}	
		
		connRedis, err := store.RedisConn()
		if err != nil {
			l.Error(err.Error())
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp{
				"success": false,
				"code": http.StatusInternalServerError,
				"msg": "Internal Server Error",
			})
			return
		}

		errRedis := connRedis.Set(ctx, userID.String(), tokens.Refresh, 0).Err()
		if errRedis != nil {
			l.Error(err.Error())
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp{
				"success": false,
				"code": http.StatusInternalServerError,
				"msg": "Internal Server Error",
			})
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": true, 
			"tokens": resp{
				"access": tokens.Access,
				"refresh": tokens.Refresh,
			},
		})
	} else {	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnauthorized,
			"msg": "unauthorized, your session was ended earlier",
		})
	}


}