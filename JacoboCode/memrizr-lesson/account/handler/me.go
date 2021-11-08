package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/memrizr/model"
	"github.com/jacobsngoodwin/memrizr/model/apperrors"
)

// Me handler calls service for getting
// a User's details
func (h *Handler) Me(c *gin.Context) {
	// A *model.User will eventually be added to context in middleware
	user, exists := c.Get("user")

	// Methods which require a valid user
	if !exists {
		log.Printf("Unable to extract user from request context for unknown user: %v\n", c)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})

		return
	}

	uid := user.(*model.User).UID

	//gin.Context satisfies go's context.Context interface
	u, err := h.UserService.Get(c, uid)

	if err != nil {
		log.Printf("Unable to find user: %v\n %v", uid, err)
		e := apperrors.NewNotFound("User", uid.String())

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})

}
