package forms

type LoginByPassWordForm struct {
	Mobile     string `form:"mobile" json:"mobile" binding:"required,mobile"`
	PassWord   string `form:"password" json:"password" binding:"required,max=20,min=3"`
	CaptchaId  string `form:"captcha_id" json:"captcha_id" binding:"required,max=20,min=3"`
	VerifyCode string `form:"verify_code" json:"verify_code" binding:"required,max=5,min=5"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	PassWord string `form:"password" json:"password" binding:"required,max=20,min=3"`
	SmsCode  string `form:"sms_code" json:"sms_code" binding:"required,max=6,min=5"`
	NickName string `form:"nick_name" json:"nick_name" binding:"required,max=30,min=5"`
}
