package handler

import (
	"net/http"
	"testcrudweb/internal/user"
	"testcrudweb/internal/util"
)

type UserHandler struct {
	Service *user.UserService
}

func (handler *UserHandler) CreateUser(responseWriter http.ResponseWriter, request *http.Request) {
	var userRequest user.CreateUserRequest

	errorOnParse := util.ParseJSONRequestBody(request, &userRequest)

	if errorOnParse != nil {
		http.Error(responseWriter, "JSON inválido: "+errorOnParse.Error(), http.StatusBadRequest)
		return
	}

	createdUser, createError := handler.Service.CreateUser(userRequest)

	if createError != nil {
		http.Error(responseWriter, "Erro ao criar usuário: "+createError.Error(), http.StatusInternalServerError)
		return
	}

	util.WriteJSONResponse(responseWriter, http.StatusCreated, createdUser)
}

func (handler *UserHandler) FindUserById(responseWriter http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")

	user, findError := handler.Service.FindUserById(id)
	if findError != nil {
		http.Error(responseWriter, "Erro ao buscar usuário: "+findError.Error(), http.StatusInternalServerError)
		return
	}

	util.WriteJSONResponse(responseWriter, http.StatusOK, user)
}
