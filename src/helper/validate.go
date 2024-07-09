package helper

import (
	"ginpet/src/model"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()

	return &Validator{validate: v}
}

func (v *Validator) ValidateClient(client model.Client) error {
	return v.validate.Struct(client)
}

func (v *Validator) ValidatePet(pet model.Pet) error {
	return v.validate.Struct(pet)
}

func (v *Validator) ValidateSP(sp model.ServiceProvider) error {
	return v.validate.Struct(sp)
}

func (v *Validator) ValidateReview(review model.Review) error {
	return v.validate.Struct(review)
}
