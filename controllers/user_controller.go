package controllers

import (
	"go-crud/config"
	"go-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {

	rows, err := config.GetDB().Query("SELECT name, is_verified, role FROM user")

	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		// http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	defer rows.Close()

	var userList []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Name, &user.IsVerified, &user.Role)
		if err != nil {
			ctx.JSON(400, gin.H{"message": err.Error()})
			// http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		userList = append(userList, user)
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{"users": userList})
	// ctx.JSON(http.StatusOK, userList)
	// json.NewEncoder(writer).Encode(userList)
}

func AddNewUser(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// err := json.NewDecoder(request.Body).Decode(&newUser)

	// if err != nil {
	// 	http.Error(writer, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// Insert the new user into the database
	id, insertErr := models.InsertUser(newUser)
	if insertErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
		// http.Error(writer, insertErr.Error(), http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully!", "id": id, "user": newUser})

	// writer.Header().Set("Content-Type", "application/json")

	// type UserResponse struct {
	// 	Message string      `json:"message"`
	// 	User    models.User `json:"user"`
	// 	Id      int         `json:"id"`
	// }

	// userResponse := UserResponse{
	// 	Message: "User Created Successfully!",
	// 	User:    newUser,
	// 	Id:      id,
	// }

	// json.NewEncoder(writer).Encode(userResponse)
}

func DeleteUser(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = models.DeleteUserByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Deleted Successfully!", "id": id})

	// params := mux.Vars(request)
	// id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	http.Error(writer, "Invalid user ID", http.StatusBadRequest)
	// 	return
	// }

	// err = models.DeleteUserByID(id)

	// if err != nil {
	// 	http.Error(writer, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// type DeleteUserResponse struct {
	// 	Message string `json:"message"`
	// 	ID      int    `json:"id"`
	// }

	// deleteUserResponse := DeleteUserResponse{Message: "User Deleted Successfully!", ID: id}

	// json.NewEncoder(writer).Encode(deleteUserResponse)
}

func UpdateUser(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser models.User
	if err := ctx.BindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateError := models.UpdateUserByID(id, updatedUser)
	if updateError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": updateError.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User Updated Successfully!", "user": updatedUser})

	// params := mux.Vars(request)

	// id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	http.Error(writer, "Invalid user ID", http.StatusBadRequest)
	// 	return
	// }

	// var updatedUser models.User

	// err = json.NewDecoder(request.Body).Decode(&updatedUser)

	// if err != nil {
	// 	http.Error(writer, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// updateError := models.UpdateUserByID(id, updatedUser)

	// if updateError != nil {
	// 	http.Error(writer, updateError.Error(), http.StatusBadRequest)
	// 	return
	// }

	// type UpdateUserResonse struct {
	// 	Message string
	// 	User    models.User
	// }

	// updateUserResonse := UpdateUserResonse{Message: "User Updated Successfully!", User: updatedUser}

	// json.NewEncoder(writer).Encode(updateUserResonse)
}
