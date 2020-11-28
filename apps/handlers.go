package apps

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSettingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"channelID": RLPConfig.ChannelID,
		"secret":    RLPConfig.ChannelSecret,
	})
}

func RequestPaymentHandler(c *gin.Context) {
	res, err := RequestPayment()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func ConfirmPaymentHandler(c *gin.Context) {
	transactionId := c.Param("transactionId")
	res, err := ConfirmPayment(transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func CapturePaymentHandler(c *gin.Context) {
	transactionId := c.Param("transactionId")
	res, err := CapturePayment(transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func RefundPaymentHandler(c *gin.Context) {
	transactionId := c.Param("transactionId")
	res, err := RefundPayment(transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}
