package helper

import (
	"encoding/json"
	"net/http"

	"github.com/komblog/restful-api-go/model/domain"
	"github.com/komblog/restful-api-go/model/web"
)

func ToUserDTO(user domain.User) web.UserDTO {
	return web.UserDTO{
		Id:         user.Id,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Birth_Date: user.Birth_Date,
	}
}

func ToUsersDTO(users []domain.User) []web.UserDTO {
	usersDTO := []web.UserDTO{}
	for _, user := range users {
		usersDTO = append(usersDTO, ToUserDTO(domain.User(user)))
	}

	return usersDTO
}

func GetRequestFromBody(request *http.Request, userRequest interface{}) {
	reader := json.NewDecoder(request.Body)
	errDecoder := reader.Decode(userRequest)
	PanicIfError(errDecoder)
}

func SendWebResponse(writer http.ResponseWriter, webResponse interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(webResponse)
	PanicIfError(errEncode)
}
