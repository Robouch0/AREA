//
// EPITECH PROJECT, 2024
// AREA
// File description:
// users
//

package controllers

import (
	"area/db"
	"area/models"
	conv_utils "area/utils/convUtils"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserInformations struct {
	ID        uint   `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`
	FirstName string `bun:"first_name" json:"first_name"`
	LastName  string `bun:"last_name" json:"last_name"`

	Email    string `bun:"email" json:"email"`
	Password string `bun:"password" json:"password"`
}

func UserModelToUserInfos(data *models.User) *UserInformations {
	return &UserInformations{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
	}
}

// Get token godoc
// @Summary      Create a new user
// @Description  Create a new user in database
// @Security ApiKeyAuth
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 userInfos body	UserInformations	true 	"User's information"
// @Success      200  {object}  UserInformations
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /user/ [post]
func CreateNewUser(userDb *db.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newUser, err := conv_utils.IoReaderToStruct[UserInformations](&r.Body)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		res, err := userDb.CreateUser(&models.User{
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
			Email:     newUser.Email,
			Password:  newUser.Password,
		})

		if err != nil {
			fmt.Printf("User Creation Error: %v\n", err)
			return
		}

		err = json.NewEncoder(w).Encode(UserModelToUserInfos(res))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}
}

// func GetAllUsers(userDb *db.UserDb) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		res, err := userDb.GetUsers()
// 		if err != nil {
// 			fmt.Printf("Error: %v\n", err)
// 			return
// 		}

// 		err = json.NewEncoder(w).Encode(res) // Convert to []UserInformations
// 		if err != nil {
// 			fmt.Printf("Error: %v\n", err)
// 			return
// 		}
// 	}
// }

// Get token godoc
// @Summary      Get User By ID
// @Description  Get user's information based on his ID
// @Security ApiKeyAuth
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  UserInformations
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /user/me [get]
func GetUserByID(userDb *db.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			return
		}

		res, err := userDb.GetUserByID(int(userID))
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 422, fmt.Sprintf("error: no such user with id: %d", userID))
			log.Printf("Error: %v\n", err)
			return
		}

		err = json.NewEncoder(w).Encode(UserModelToUserInfos(res))
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 422, fmt.Sprintf("Error: %v\n", err))
			log.Printf("Error: %v\n", err)
			return
		}
	}
}

// Get token godoc
// @Summary      Update User datas
// @Description  Update some informations about the user
// @Security ApiKeyAuth
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 updatableDatas body	models.UpdatableUserData	true 	"Updatable user's informations"
// @Success      200  {object}  models.UpdatableUserData
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /user/me [put]
func UpdateUserDatas(userDb *db.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			log.Println(err)
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			return
		}

		updateData, err := conv_utils.IoReaderToStruct[models.UpdatableUserData](&r.Body)
		if err != nil {
			log.Println(err)
			http_utils.WriteHTTPResponseErr(&w, 422, fmt.Sprintf("Invalid request format"))
			return
		}

		_, err = userDb.UpdateUserData(int(userID), updateData)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 422, fmt.Sprintf("Invalid informations sent to update user"))
			log.Println(err)
			return
		}
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(updateData)
	}
}
