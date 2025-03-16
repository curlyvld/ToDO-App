package handler

import "github.com/gin-gonic/gin"

func (h *Handler) createList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
