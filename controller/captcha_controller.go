package controller

import (
	"captcha/model"
	"captcha/recaptcha"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var GetCaptcha = func(context *gin.Context) {
	baseResponse := model.NewBaseResponse()
	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}
	if d.CaptchaId != "" {
		baseResponse.GetSuccessResponse()
		var captcha model.CaptchaResponse
		captcha.CaptchaId = d.CaptchaId
		captcha.ImageUrl = "/show/" + d.CaptchaId + ".png"
		baseResponse.Data = captcha
	} else {
		baseResponse.GetFailureResponse(model.SYSTEM_ERROE)
	}
	context.JSON(http.StatusOK, baseResponse)
}

var VerifyCaptcha = func(context *gin.Context) {
	baseResponse := model.NewBaseResponse()
	captchaId := context.Request.FormValue("captchaId")
	value := context.Request.FormValue("value")
	if captchaId == "" || value == "" {
		baseResponse.GetFailureResponse(model.QUERY_PARAM_ERROR)
	} else {
		if captcha.VerifyString(captchaId, value) {
			baseResponse.GetSuccessResponse()
			baseResponse.Message = "验证成功"
		} else {
			baseResponse.GetFailureResponse(model.CAPTCHA_ERROR)
		}
	}
	context.JSON(http.StatusOK, baseResponse)
}

var GetCaptchaPng = func(context *gin.Context) {
	source := context.Param("source")
	logrus.Info("GetCaptchaPng : " + source)
	recaptcha.ServeHTTP(context.Writer, context.Request)
}
