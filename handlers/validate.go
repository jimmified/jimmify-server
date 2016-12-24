package handlers

import (
	"errors"
	"jimmify-server/db"
	"log"

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
		//failed validation
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		return errors.New("Invalid Query")
	}

	//check for allowed type
	if !types[q.Type] {
		return errors.New("Unknown Type")
	}

	return nil
}
