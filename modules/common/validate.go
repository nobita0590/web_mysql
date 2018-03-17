package common

import (
	"gopkg.in/kataras/iris.v8"
	"gopkg.in/go-playground/validator.v9"

	"github.com/go-playground/locales/en"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	"github.com/go-playground/universal-translator"
	"fmt"
	"github.com/go-playground/form"
	"strings"
)

func ReadStruct(context iris.Context,target interface{},isValid bool,ignoreFields ...string) (err error,valid validator.ValidationErrorsTranslations) {
	decoder := form.NewDecoder()
	if  err = decoder.Decode(target,context.FormValues());err == nil && isValid{
		err,valid = ValidateStruct(context,target,ignoreFields...)
	}
	return
}

func ReadJSONStruct(context iris.Context,target interface{},isValid bool,ignoreFields ...string) (err error,valid validator.ValidationErrorsTranslations) {
	if  err = context.ReadJSON(target);err == nil && isValid{
		err,valid = ValidateStruct(context,target,ignoreFields...)
	}
	return
}

func ValidateStruct(context iris.Context,target interface{},ignoreFields ...string) (error,validator.ValidationErrorsTranslations) {
	validate := validator.New()
	enLang := en.New()
	uni := ut.New(enLang)

	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.StructExcept(target,ignoreFields...)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		tran := errs.Translate(trans)
		return err,convertValidation(tran)
	}
	return nil,nil
}
func convertValidation(validMess validator.ValidationErrorsTranslations) validator.ValidationErrorsTranslations {
	if(validMess == nil){
		return validMess
	}
	newMess := make(validator.ValidationErrorsTranslations)
	for k,mess := range validMess{
		if index := strings.Index(k,".");index != -1{
			k = k[index+1:]
		}
		newMess[k] = mess
	}

	return newMess
}

func parserValidate(err error)  {
	if val, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println(err,val)
		return
	}

	for _, err := range err.(validator.ValidationErrors) {

		fmt.Println(err.Namespace())
		fmt.Println(err.Field())
		fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
		fmt.Println(err.StructField())     // by passing alt name to ReportError like below
		fmt.Println(err.Tag())
		fmt.Println(err.ActualTag())
		fmt.Println(err.Kind())
		fmt.Println(err.Type())
		fmt.Println(err.Value())
		fmt.Println(err.Param())
		fmt.Println()
	}
}