package types

type OTPReqBody struct {
	Id    int64
	Email string
}

type VerifyOtpCode struct {
	OtpCode int64
}
