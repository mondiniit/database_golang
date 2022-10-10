package game_application

import (
	"net/http"
	"strings"

	my_postgres "database_golang/src/database"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Input string `json:"input"`
}

func ResponseGame(c *gin.Context) {
	var input Result

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BAD REQUEST"})
		return
	}
	if s := strings.Trim(input.Input, ","); len(s) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "INVALID INPUT"})
		return
	}
	s := strings.Replace(input.Input, " ", "+", len(input.Input))
	s = strings.Trim(s, ",")
	ss := strings.Split(s, ",")

	if err := ValidateArguments(ss); err == nil {
		numbers, operators, err := InitGameStructure(ss)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "INVALID INPUT"})
			return
		}
		r, err := Calculate(numbers, operators)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "CAN'T DIVIDE BY ZERO"})
		} else if r == 42 {
			try := strings.Join(ss, "")
			s := GetHints(try, my_postgres.Connection())
			if s == "CCCCCC" {
				c.JSON(http.StatusOK, gin.H{"content": s})
			} else {
				c.JSON(http.StatusOK, gin.H{"content": s})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "EQUATION MUST RESULT IN 42"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "INVALID INPUT"})
	}
}
