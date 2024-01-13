package handler

import (
	"fitter-srv/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createStatistic(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	setId, err := strconv.Atoi(c.Param("idS"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid set's id param")
		return
	}

	var input model.Statistic
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Statistic.Create(userId, setId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllStatistics(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	setId, err := strconv.Atoi(c.Param("idS"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid set's id param")
		return
	}
	statistics, err := h.services.Statistic.GetAll(userId, setId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statistics)
}

func (h *Handler) getStatisticByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	statisticId, err := strconv.Atoi(c.Param("idSs"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid statistic's id param")
		return
	}
	statistic, err := h.services.Statistic.GetById(userId, statisticId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statistic)
}

func (h *Handler) updateStatistic(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("idSs"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid statistic's id param")
		return
	}

	var input model.UpdateStatisticInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Statistic.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) deleteStatistic(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	statisticId, err := strconv.Atoi(c.Param("idSs"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid statistic's id param")
		return
	}
	err = h.services.Statistic.Delete(userId, statisticId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
