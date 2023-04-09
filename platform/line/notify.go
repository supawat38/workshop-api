package notify

import (
	"github.com/juunini/simple-go-line-notify/notify"
)

var (
	accessToken = "6oKwJl1hz74n89J9zYduXSTvxMEOb9Q2oZri59yZKxq"
)

// module default register
func LineSendOTP(otp string, refno string, module string) (success bool) {

	if otp == "" || refno == "" {
		success = false
		return
	}

	if module == "" {
		module = "register"
	}

	var txtModule = ""

	switch module {
	case "register":
		txtModule = "สมัครสมาชิก"
	}

	var message = "รหัสผ่าน OTP " + otp + " ( Ref No. " + refno + " ) สำหรับ " + txtModule + " GameMarket ใช้ได้ภายใน 5 นาที ห้ามบอกรหัส OTP กับบุคคลอื่นไม่ว่ากรณีใดๆ"

	if err := notify.SendText(accessToken, message); err != nil {
		success = false
		return
		// panic(err)
	}

	success = true

	return

}
