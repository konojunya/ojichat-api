package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/greymd/ojichat/generator"
)

func responseBadRequest(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(400, gin.H{
		"error": msg,
	})
}

func handleReponse(c *gin.Context) {

	targetName := c.DefaultQuery("name", "")
	emojiNumStr := c.DefaultQuery("e", "4")
	punctuationLabelStr := c.DefaultQuery("p", "0")

	emojiNum, err := strconv.Atoi(emojiNumStr)
	if err != nil {
		responseBadRequest(c, "It can only handle numbers.")
		return
	}

	punctuationLabel, err := strconv.Atoi(punctuationLabelStr)
	if err != nil {
		responseBadRequest(c, "It can only handle numbers.")
		return
	}
	if punctuationLabel > 3 || punctuationLabel < 0 {
		responseBadRequest(c, "Possible values for punctuation label are 0 to 3")
		return
	}

	config := generator.Config{
		TargetName:        targetName,
		EmojiNum:          emojiNum,
		PunctiuationLebel: punctuationLabel,
	}
	result, err := generator.Start(config)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, gin.H{
		"message": result,
	})
}

func main() {
	r := gin.Default()
	r.GET("/", handleReponse)
	r.Run()
}
