package types

type OTPReqBody struct {
	Id    int64
	Email string
}

type VerifyOtpCode struct {
	UserId  int64  `json:"user_id"`
	OtpCode string `json:"otp_code"`
}
