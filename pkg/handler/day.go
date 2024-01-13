package handler

import (
	"fitter-srv/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createDay(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	programId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid program id param")
		return
	}

	var input model.TrainDay
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Day.Create(userId, programId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllDays(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	programId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid program id param")
		return
	}
	days, err := h.services.Day.GetAll(userId, programId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, days)
}

func (h *Handler) getDayByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	dayId, err := strconv.Atoi(c.Param("idD"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid day id param")
		return
	}
	day, err := h.services.Day.GetById(userId, dayId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, day)
}

func (h *Handler) updateDay(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("idD"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid day id param")
		return
	}

	var input model.UpdateDayInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Day.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) deleteDay(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	dayId, err := strconv.Atoi(c.Param("idD"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid day id param")
		return
	}
	err = h.services.Day.Delete(userId, dayId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
