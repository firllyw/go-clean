package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	member "github.com/goclean/internal/model"
	"github.com/labstack/echo"
)

type MemberHandler struct {
	MemberUseCase member.Service
}

func NewMemberHandler(e *echo.Echo, svc member.Service) {
	handler := &MemberHandler{
		MemberUseCase: svc,
	}

	e.GET("/member", handler.Index)
}

func (m *MemberHandler) Index(c echo.Context) error {
	pageS, sizeS := c.QueryParam("page"), c.QueryParam("size")
	page, _ := strconv.Atoi(pageS)
	size, _ := strconv.Atoi(sizeS)
	members, err := m.MemberUseCase.Index(c.Request().Context(), map[string]interface{}{}, page, size)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	return c.JSON(http.StatusOK, members)
}

func (m *MemberHandler) Insert(c echo.Context) error {
	data := c.Request().Body
	memberStruct := member.Member{}
	err := json.NewDecoder(data).Decode(&memberStruct)
	if err != nil {
		return c.JSON(500, err.Error())
	}
	member, err := m.MemberUseCase.Insert(c.Request().Context(), memberStruct)
	return c.JSON(202, member)
}
