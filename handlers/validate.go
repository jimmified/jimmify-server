package handlers

import (
	"errors"
	"jimmify-server/auth"
	"jimmify-server/db"

	"github.com/astaxie/beego/validation"
)

//validateQuery validate field types and length
func validateQuery(q db.Query) error {
	valid := validation.Validation{}
	types := map[string]bool{
		"search": true,
	}
	//create rules
	valid.Required(q.Text, "text")         //require the text field
	valid.MaxSize(q.Text, 255, "textSize") //verify it is not too long
	valid.Required(q.Type, "textType")     //require the type field
	valid.MaxSize(q.Type, 20, "typeSize")  //verify it is not too long
	valid.Alpha(q.Type, "typeType")        //verify type
	if valid.HasErrors() {
		return errors.New("Invalid Query") //failed validation
	}
	//check for allowed type
	if !types[q.Type] {
		return errors.New("Unknown Type")
	}
	return nil
}

//validateAnswer check for valid answer
func validateAnswer(q db.Query) error {
	valid := validation.Validation{}
	//create rules
	valid.Required(q.Key, "key")
	valid.MaxSize(q.Answer, 800, "answerSize")
	if valid.HasErrors() {
		return errors.New("Invalid Answer") //failed validation
	}
	return nil
}

//validateAnswer check for valid answer
func validateCheck(q db.Query) error {
	valid := validation.Validation{}
	//create rules
	valid.Required(q.Key, "key")
	if valid.HasErrors() {
		return errors.New("No key") //failed validation
	}
	return nil
}

//validateLogin check login credentials
func validateLogin(c Credentials) error {
	if c.Password != auth.JPASS {
		return errors.New("Invalid Password")
	}
	return nil
}
