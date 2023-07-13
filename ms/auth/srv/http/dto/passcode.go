package dto

import "time"

type PasscodeFinishBody struct {
	Id   string `json:"id" validate:"required,uuid4"`
	Code string `json:"code" validate:"required"`
}

type PasscodeInitBody struct {
	UserId  string  `json:"user_id" validate:"required,uuid4"`
	EmailId *string `json:"email_id"`
}

type PasscodeReturn struct {
	Id        string    `json:"id"`
	TTL       int32     `json:"ttl"`
	CreatedAt time.Time `json:"created_at"`
}
