package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/komblog/restful-api-go/helper"
	"github.com/komblog/restful-api-go/model/web"
	"github.com/komblog/restful-api-go/service"
)

type UserController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := &web.CreateUserRequest{}
	helper.GetRequestFromBody(request, userRequest)

	userResult := controller.service.Create(request.Context(), *userRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResult,
	}
	helper.SendWebResponse(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userID := params.ByName("userId")
	userId, errAtoi := strconv.Atoi(userID)
	helper.PanicIfError(errAtoi)

	userResult := controller.service.FindById(request.Context(), userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResult,
	}

	helper.SendWebResponse(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usersResult := controller.service.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   usersResult,
	}

	helper.SendWebResponse(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, errAtoi := strconv.Atoi(userId)
	helper.PanicIfError(errAtoi)

	userRequest := web.CreateUserRequest{}
	userRequest.Id = id

	helper.GetRequestFromBody(request, &userRequest)

	userResult := controller.service.Update(request.Context(), userRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResult,
	}

	helper.SendWebResponse(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, errAtoi := strconv.Atoi(userId)
	helper.PanicIfError(errAtoi)

	findUser := controller.service.FindById(request.Context(), id)

	controller.service.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   findUser,
	}

	helper.SendWebResponse(writer, webResponse)
}
