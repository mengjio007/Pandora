package Pinocchio

import (
	"github.com/dchest/captcha"
	"io"
)

func CaptchaNew() string {
	return captcha.New()
}

func CaptchaReload(captchaId string) bool {
	return captcha.Reload(captchaId)
}

func CaptchaWriteImage(w io.Writer, id string, width, height int) error {
	return captcha.WriteImage(w, id, width, height)
}

func CaptchaVerify(id string, digits string) bool {
	return captcha.VerifyString(id, digits)
}
