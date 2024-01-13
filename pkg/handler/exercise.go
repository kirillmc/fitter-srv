package handler

import (
	"fitter-srv/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createExercise(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	dayId, err := strconv.Atoi(c.Param("idD"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid day id param")
		return
	}

	var input model.Exercise
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Exercise.Create(userId, dayId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllExercises(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	dayId, err := strconv.Atoi(c.Param("idD"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid day id param")
		return
	}
	exercises, err := h.services.Exercise.GetAll(userId, dayId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, exercises)
}

func (h *Handler) getExerciseByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	exerciseId, err := strconv.Atoi(c.Param("idE"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid exercise id param")
		return
	}
	day, err := h.services.Exercise.GetById(userId, exerciseId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, day)
}

func (h *Handler) updateExercise(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("idE"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input model.UpdateExerciseInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Exercise.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) deleteExercise(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	exerciseId, err := strconv.Atoi(c.Param("idE"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid exercise id param")
		return
	}
	err = h.services.Day.Delete(userId, exerciseId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
