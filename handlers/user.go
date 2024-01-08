package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/amosehiguese/restaurant-api/models"
	"github.com/amosehiguese/restaurant-api/store"
	"github.com/amosehiguese/restaurant-api/types"
	"github.com/google/uuid"
)

// GetUsers returns all users
// @Summary List all users
// @Description Get all users stored in the database
// @Tags User
// @Produce json
// @Router /users [get]
// @Success 200 {object} models.User
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
func GetUsers(w http.ResponseWriter, r *http.Request) {
	s, e, err := paginate(w, r)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return 
	}

	q := store.GetQuery()

	result, err := q.GetAllUsers(ctx)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	if *e < len(result) && len(result[*s:*e]) == pageSize {
		result = result[*s:*e]
	} else if *e >= len(result) && *s < len(result) {
		result = result[*s:]
	} else if *e >= len(result) && *s >= len(result) && result != nil {
		*s = 0
		*e = pageSize
		result = result[*s:*e]
	}	


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": result,
	})
}

// RetrieveUser renders the user with the given id 
// @Summary Get user by id
// @Description RetrieveUser returns a single user by id
// @Tags User
// @Produce json
// @Param id path string true "user id"
// @Router /users/{id} [get]
// @Success 200 {object} models.User
// @Failure 400 {object} http.StatusBadRequest
// @Failure 404 {object} http.StatusNotFound
func RetrieveUser(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	userID, err := uuid.Parse(id)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}

	q := store.GetQuery()
	user, err := q.RetrieveUser(ctx, userID)
	if err != nil {
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusNotFound,
			"msg": fmt.Sprintf("user with this ID %v not found", userID),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"user_id": user,
	})	
}

// UpdateUser modifies the user with the given id 
// @Summary Modify user by id
// @Description UpdateUser modifies a single menu by id
// @Tags User
// @Produce json
// @Param id path string true "user id"
// @Router /users/{id} [patch]
// @Success 200 {object} models.User
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	userID, err := uuid.Parse(id)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}

	var payload types.UserPayload
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		l.Errorln(err)
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnprocessableEntity,
			"msg": "Unprocessable entity",
		})
		return
	}

	v := types.NewValidator()

	if err := v.Struct(payload); err != nil{
		json.NewEncoder(w).Encode(resp{
			"error": true,
			"code": http.StatusBadRequest,
			"msg":types.ValidatorErrors(err),
		})
		return
	}

	q := store.GetQuery()

	user := &models.UpdateUserParams{
		ID: userID,
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Username: payload.Username,
		
	}

	err = q.UpdateUser(ctx, *user)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	str := fmt.Sprintf("Successfully update user with id %s", userID)
	json.NewEncoder(w).Encode(resp{
		"success":true,
		"msg": str,
	})
}

// CreateRole writes a role to the database
// @Summary Creates a role
// @Description Creates a role in the database
// @Tags Role
// @Produce json
// @Router /roles [post]
// @Success 200 {object} models.Role
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
func CreateRole(w http.ResponseWriter, r *http.Request) {
	var role models.CreateRoleParams
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		l.Errorln(err)
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnprocessableEntity,
			"msg": "Unprocessable entity",
		})
		return
	}

	v := types.NewValidator()
	if err := v.Struct(role); err != nil {
		json.NewEncoder(w).Encode(resp{
			"error": true,
			"code": http.StatusBadRequest,
			"msg":types.ValidatorErrors(err),
		})
		return
	}

	q := store.GetQuery()

	newRole, err := q.CreateRole(ctx, role)
		if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"role": newRole,
	})
}

// GetRoles returns all roles
// @Summary List all roles
// @Description Get all roles stored in the database
// @Tags Role
// @Produce json
// @Router /roles [get]
// @Success 200 {object} models.Role
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
func GetRoles(w http.ResponseWriter, r *http.Request) {
	s, e, err := paginate(w, r)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return 
	}	

	q := store.GetQuery()

	result, err := q.GetAllRoles(ctx)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	if *e < len(result) && len(result[*s:*e]) == pageSize {
		result = result[*s:*e]
	} else if *e >= len(result) && *s < len(result) {
		result = result[*s:]
	} else {
		*s = 0
		*e = pageSize
		result = result[*s:*e]
	}	

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"data": result,
	})
}

// UpdateRole modifies the role with the given id 
// @Summary Modify role by id
// @Description UpdateRole modifies a single role by id
// @Tags Role
// @Produce json
// @Param id path string true "role id"
// @Router /roles/{id} [patch]
// @Success 200 {object} models.Role
// @Failure 400 {object} http.StatusBadRequest
// @Failure 422 {object} http.StatusUnprocessableEntity
// @Failure 500 {object} http.StatusInternalServerError
func UpdateRole(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	roleId, err := strconv.Atoi(id)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}

	var payload types.RolePayload
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		l.Errorln(err)
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusUnprocessableEntity,
			"msg": "Unprocessable entity",
		})
		return
	}

	v := types.NewValidator()

	if err := v.Struct(payload); err != nil{
		json.NewEncoder(w).Encode(resp{
			"error": true,
			"code": http.StatusBadRequest,
			"msg":types.ValidatorErrors(err),
		})
		return
	}

	role := &models.UpdateRoleParams{
		ID: int32(roleId),
		Name: payload.Name,
		Description: payload.Description,
	}

	q := store.GetQuery()
	err = q.UpdateRole(ctx, *role)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	str := fmt.Sprintf("Successfully update role with id %v", roleId)
	json.NewEncoder(w).Encode(resp{
		"success":true,
		"msg": str,
	})
}


// DeleteRole remove the role with the given id 
// @Summary Removes role by id
// @Description Removes a single role by id from the database
// @Tags Role
// @Produce json
// @Param id path string true "role id"
// @Router /roles/{id} [delete]
// @Success 200 {object} string
// @Failure 400 {object} http.StatusBadRequest
// @Failure 500 {object} http.StatusInternalServerError
func DeleteRole(w http.ResponseWriter, r *http.Request) {
	id := getField(r, "id")
	roleID, err := strconv.Atoi(id)
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusBadRequest,
			"msg": "Bad request",
		})
		return
	}
	q := store.GetQuery()

	err = q.DeleteRole(ctx, int32(roleID))
	if err != nil {
		l.Error(err.Error())
		json.NewEncoder(w).Encode(resp{
			"success": false,
			"code": http.StatusInternalServerError,
			"msg": "Internal server error",
		})
		return
	}

	dataResp := fmt.Sprintf("Role with id %s is successfully deleted", id)
	json.NewEncoder(w).Encode(resp{
		"success": true,
		"msg": dataResp,
	})
}