package rest

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"user-service/internal/application/interfaces"
	"user-service/internal/interface/api/rest/dto/mapper"
	"user-service/internal/interface/api/rest/dto/request"
	"user-service/internal/validator"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	accountsService interfaces.AccountsService
}

func NewUserController(e *echo.Echo, accountsService interfaces.AccountsService) *UserController {
	controller := &UserController{
		accountsService,
	}

	router := e.Group("/api/v1/user")

	router.POST("/signup", controller.RegisterUserController)

	return controller
}

// @Summary Create a new user
// @Description Sign up a user
// @Tags Accounts Service
// @Accept json
// @Produce json
// @Param request body request.RegisterUserRequest true "User registeration request"
// @Success 201 {object} response.RegisterUserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} resterror.ErrInternal
// @Router /user/signup [post]
func (ac *UserController) RegisterUserController(c echo.Context) error {
	var registerUserRequest request.RegisterUserRequest

	if err := c.Bind(&registerUserRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	registerUserCommand, err := registerUserRequest.ToRegisterUserCommand()
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			return echo.NewHTTPError(http.StatusBadRequest, errors.ToMap())
		}
		return echo.ErrInternalServerError
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := ac.accountsService.Register(ctx, registerUserCommand)
	if err != nil {
		fmt.Println("error: ", err)
		return echo.ErrInternalServerError
	}

	response := mapper.ToRegisterUserResponse(result.Result)

	return c.JSON(http.StatusCreated, response)
}
