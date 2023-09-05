package uservalidator

import (
	"Q/A-GameApp/dto"
	"Q/A-GameApp/pkg/errmsg"
	"Q/A-GameApp/pkg/richerror"
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}
type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}
func (v Validator) ValidateRegisterRequest(req dto.RegisterRequest) (error, map[string]string) {
	const op = "UserValidator.validateRegisterRequest"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Password, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-z0-9!@#$%&*]{8,}`))),
		validation.Field(&req.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile(`^09[0-9]{9}$`)),
			validation.By(v.checkPhoneNumberUniqueness)),
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

		return richerror.New(op).WhitMessage(errmsg.ErrorMsgInvalidInput).WhitKind(richerror.KindInvalid).
			WhitMeta(map[string]interface{}{"request:": req}).WhitWarpError(err), fieldErr
	}

	return nil, nil

}
func (v Validator) checkPhoneNumberUniqueness(value interface{}) error {
	phoneNumber := value.(string)
	if isUnique, err := v.repo.IsPhoneNumberUnique(phoneNumber); err != nil || !isUnique {
		if err != nil {
			return err
		}
		if !isUnique {
			return fmt.Errorf(errmsg.ErrorMsgPhoneNumberIsNotUnique)
		}
	}
	return nil
}
