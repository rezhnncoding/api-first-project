package controller

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"puppy/Utility"
	httpResponse "puppy/ViewModel/common/httpresponse"
	userViewModel "puppy/ViewModel/user"
	"puppy/service"
)

type UserController interface {
	GetUserList(c echo.Context) error
	CreateNewUser(c echo.Context) error
	EditUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	EditUserRole(c echo.Context) error
	EditUserPassword(c echo.Context) error
	UploadAvatar(c echo.Context) error
}

type userController struct {
}

func NewUserController() UserController {
	return userController{}
}

func (uc userController) GetUserList(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	fmt.Println(apiContext.GetUserId())

	userService := service.NewUserService()
	userList, err := userService.GetUserList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userList))
}
func (uc userController) CreateNewUser(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	operatorUserId, err := apiContext.GetUserId()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	userService := service.NewUserService()
	//isValid := userService.IsUserValidForAccess(operatorUserId, "CreateUser")
	//if !isValid {
	//	return c.JSON(http.StatusForbidden, "")
	//}

	newUser := new(userViewModel.CreateNewUserViewModel)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("Data not found"))
	}

	if err := c.Validate(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse(err))
	}

	file, err := apiContext.FormFile("file")
	if err == nil {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		wd, err := os.Getwd()
		imageServerPath := filepath.Join(wd, "wwwroot", "images", "userAvatar", "dcfvgbhnjm"+filepath.Ext(file.Filename))

		des, err := os.Create(imageServerPath)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer des.Close()

		_, err = io.Copy(des, src)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		newUser.AvatarName = "dcfvgbhnjm" + filepath.Ext(file.Filename)
	}

	newUser.CreatorUserId = operatorUserId

	newUserId, err := userService.CreateNewUser(*newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: newUserId,
	}
	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userResData))
}

func (uc userController) EditUser(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()
	newUserData := new(userViewModel.EditUserViewModel)

	if err := c.Bind(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	newUserData.TargetUserId = targetUserId

	if !userService.IsUserExist(targetUserId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := userService.EditUser(*newUserData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}

func (uc userController) DeleteUser(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()

	if !userService.IsUserExist(targetUserId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := userService.DeleteUser(targetUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}

func (uc userController) EditUserRole(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()
	newUserData := new(userViewModel.EditUserRoleViewModel)

	if err := c.Bind(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	newUserData.TargetUserId = targetUserId

	if !userService.IsUserExist(targetUserId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := userService.EditUserRole(*newUserData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}

func (uc userController) EditUserPassword(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()
	newUserData := new(userViewModel.EditUserPasswordViewModel)

	if err := c.Bind(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	newUserData.TargetUserId = targetUserId

	if !userService.IsUserExist(targetUserId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := userService.EditUserPassword(*newUserData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}

func (uc userController) UploadAvatar(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)

	file, err := apiContext.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	wd, err := os.Getwd()
	imageServerPath := filepath.Join(wd, "wwwroot", "images", "userAvatar", file.Filename)

	des, err := os.Create(imageServerPath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer des.Close()

	_, err = io.Copy(des, src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}
