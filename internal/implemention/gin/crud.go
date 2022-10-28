package ginrouter

import (
	formatter "github.com/EgorMatirov/microtasks/internal/implemention/json.formatter"
	"github.com/gin-gonic/gin"
	"net/http"
	"unicode/utf8"
)

func (rH RouterHandler) sayHello(c *gin.Context) {
	name := c.Param("name")

	if utf8.RuneCountInString(name) == 0 {
		c.JSON(http.StatusBadRequest, formatter.NewSayHelloResponse(name))

		return
	}

	c.JSON(http.StatusOK, formatter.NewSayHelloResponse(name))
	return
}
