package apps

import (
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GenerateRouter() *gin.Engine {
	gin.SetMode(GinMode)

	router := gin.New()

	httpLog := logrus.New()
	if env := os.Getenv("GO_ENV"); env == "test" {
		httpLog.SetOutput(ioutil.Discard)
	} else {
		httpLog.SetOutput(os.Stdout)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	router.GET("/", GetSettingHandler)
	router.GET("/request", RequestPaymentHandler)
	router.GET("txn/:transactionId/confirm", ConfirmPaymentHandler)
	router.GET("txn/:transactionId/capture", CapturePaymentHandler)
	router.GET("txn/:transactionId/refund", RefundPaymentHandler)

	return router
}
