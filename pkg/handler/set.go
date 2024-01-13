package handler

import (
	"fitter-srv/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createSet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	exerciseId, err := strconv.Atoi(c.Param("idE"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid exercise id param")
		return
	}

	var input model.Set
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Set.Create(userId, exerciseId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllSets(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	exerciseId, err := strconv.Atoi(c.Param("idE"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid exercise id param")
		return
	}
	sets, err := h.services.Set.GetAll(userId, exerciseId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, sets)
}

func (h *Handler) getSetByID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	setId, err := strconv.Atoi(c.Param("idS"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid set id param")
		return
	}
	day, err := h.services.Day.GetById(userId, setId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, day)
}

func (h *Handler) updateSet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("idS"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid set id param")
		return
	}

	var input model.UpdateSetInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Set.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) deleteSet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	setId, err := strconv.Atoi(c.Param("idS"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid set id param")
		return
	}
	err = h.services.Set.Delete(userId, setId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
