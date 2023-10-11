package uservalidator

import (
	"Q/A-GameApp/dto"
	"Q/A-GameApp/pkg/errmsg"
	"Q/A-GameApp/pkg/richerror"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func (v Validator) ValidateLoginRequest(req dto.LoginRequest) (map[string]string, error) {
	const op = "UserValidator.validateLoginRequest"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errmsg.ErrorMsgPhoneNumberIsNotValid),
			validation.By(v.doesPhoneNumberExist)),
		validation.Field(&req.Password, validation.Required),
	); err != nil {

		fieldErr := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErr[key] = value.Error()

				}

			}
		}

		return fieldErr, richerror.New(op).WhitMessage(errmsg.ErrorMsgInvalidInput).WhitKind(richerror.KindInvalid).
			WhitMeta(map[string]interface{}{"request:": req}).WhitWarpError(err)
	}

	return nil, nil

}
func (v Validator) doesPhoneNumberExist(value interface{}) error {
	phoneNumber := value.(string)
	_, err := v.repo.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return fmt.Errorf(errmsg.ErrorMsgNotFound)
	}

	return nil
}
