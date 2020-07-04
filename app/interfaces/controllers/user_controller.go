package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"example.com/m/usecase"
	"example.com/m/interfaces/database"
	"example.com/m/domain"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor {
			UserRepository: &database.UserRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Index(c echo.Context) error {
	users, err := controller.Interactor.Users()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (controller *UserController) Show(c echo.Context) error {
	id, idErr := strconv.Atoi(c.Param("id"))

	if idErr != nil {
		return c.String(http.StatusBadRequest, "ID Invalid!")
	}

	user, err := controller.Interactor.User(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Create(c echo.Context) error {
	u := new(domain.User)

	if bindError := c.Bind(u); bindError != nil {
		return bindError
	}

	if err := u.UserValidation(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, createError := controller.Interactor.Create(u)

	if createError != nil {
		return createError
	}
	
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Update(c echo.Context) error {
	id, idErr := strconv.Atoi(c.Param("id"))

	if idErr != nil {
		return c.String(http.StatusBadRequest, "ID Invalid!")
	}

	u := new(domain.User)
	if bindError := c.Bind(u); bindError != nil {
		return bindError
	}

	if err := u.UserValidation(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, updateError := controller.Interactor.Update(id, u)

	if updateError != nil {
		return updateError
	}
	
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Destroy(c echo.Context) error {
	id, idErr := strconv.Atoi(c.Param("id"))

	if idErr != nil {
		return c.String(http.StatusBadRequest, "ID Invalid!")
	}

	user, destroyErr := controller.Interactor.Destroy(id)

	if destroyErr != nil {
		return destroyErr
	}
	
	return c.JSON(http.StatusOK, user)
}