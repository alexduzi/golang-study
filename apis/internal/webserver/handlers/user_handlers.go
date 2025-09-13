package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alexduzi/golang-study/apis/internal/dto"
	"github.com/alexduzi/golang-study/apis/internal/entity"
	"github.com/alexduzi/golang-study/apis/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       userDB,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

// GetJWT user godoc
// @Summary           Get a user JWT
// @Description       Get a user JWT
// @Tags              users
// @Accept            json
// @Produce           json
// @Param             request      body     dto.GetJwtInput    true    "user credentials"
// @Success           201          {object} dto.GetJwtOutput
// @Failure           404          {object} Error
// @Failure           500          {object} Error
// @Router            /users/token [post]
func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var jwtInput dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&jwtInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	user, err := h.UserDB.FindByEmail(jwtInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if !user.ValidatePassword(jwtInput.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})

	accessToken := dto.GetJwtOutput{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary           Create user
// @Description       Create user
// @Tags              users
// @Accept            json
// @Produce           json
// @Param             request     body     dto.CreateUserInput   true    "user request"
// @Success           201
// @Failure           500         {object} Error
// @Router            /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	emailExists, err := h.UserDB.FindByEmail(user.Email)
	if emailExists != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	newUser, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
