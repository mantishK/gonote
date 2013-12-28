package controllers
import (
"github.com/robfig/revel"
_ "github.com/go-sql-driver/mysql"
"fmt"
"gonote/app/model"
"gonote/app/database"
"gonote/app/error"
)

type Note struct {
	*revel.Controller
}

func (c Note) GetNotes() revel.Result {
	dbMap := database.NewConnection()
	defer dbMap.Db.Close()
	notes, count, err := model.GetNotes(dbMap)
	if(err != nil) {
		return c.RenderJson(error.Error{Id:1,ErrorMessage:fmt.Sprint(err),DisplayMessage:"Some error occured"})
	}
	result := make(map[string]interface{})
	result["count"] = count
	result["notes"] = notes
	result["response"] = "ok"
	return c.RenderJson(result)
}
func (c Note) GetDetails() revel.Result {
	dbMap := database.NewConnection()
	defer dbMap.Db.Close()
	var note_id int
	c.Params.Bind(&note_id,"note_id")
	c.Validation.Required(note_id)

	//build errors
	errors := c.checkErrors()

	//check and return if errors
	if (errors != nil) {
		return c.RenderJson(errors)
	}
	note, err := model.GetDetails(dbMap,note_id)
	if(err != nil) {
		return c.RenderJson(error.Error{Id:1,ErrorMessage:fmt.Sprint(err),DisplayMessage:"Some error occured"})
	}
	result := make(map[string]interface{})
	result["note_id"] = note_id
	result["note"] = note
	result["response"] = "ok"
	return c.RenderJson(result)
}
func (c Note) Add() revel.Result {	
	dbMap := database.NewConnection()
	defer dbMap.Db.Close()
	note := model.Note{}
	c.Params.Bind(&note.Title,"title")
	c.Params.Bind(&note.Content,"content")

	//validation
	c.Validation.Required(note.Title)
	c.Validation.Required(note.Content)

	//build errors
	errors := c.checkErrors()

	//check and return if errors
	if (errors != nil) {
		return c.RenderJson(errors)
	}
	
	err := note.Save(dbMap)
	if(err != nil) {
		return c.RenderJson(error.Error{Id:1,ErrorMessage:fmt.Sprint(err),DisplayMessage:"Some error occured"})
	}
	result := make(map[string]interface{})
	result["response"] = "ok"
	result["note"] = note
	return c.RenderJson(result)
}

func (c Note) Update() revel.Result {
	dbMap := database.NewConnection()
	defer dbMap.Db.Close()
	note := model.Note{}
	c.Params.Bind(&note.Note_id,"note_id")
	c.Params.Bind(&note.Title,"title")
	c.Params.Bind(&note.Content,"content")

	//validation
	c.Validation.Required(note.Note_id)
	c.Validation.Required(note.Title)
	c.Validation.Required(note.Content)

	//build errors
	errors := c.checkErrors()

	//check and return if errors
	if (errors != nil) {
		return c.RenderJson(errors)
	}

	_,err := note.Update(dbMap)
	if(err != nil) {
		return c.RenderJson(error.Error{Id:1,ErrorMessage:fmt.Sprint(err),DisplayMessage:"Some error occured"})
	}
	result := make(map[string]interface{})
	result["response"] = "ok"
	result["note"] = note
	return c.RenderJson(result)
}
func (c Note) Delete() revel.Result {
	dbMap := database.NewConnection()
	defer dbMap.Db.Close()
	note := model.Note{}
	c.Params.Bind(&note.Note_id,"note_id")

	//validation
	c.Validation.Required(note.Note_id)

	//build errors
	errors := c.checkErrors()

	//check and return if errors
	if (errors != nil) {
		return c.RenderJson(errors)
	}


	count,err := note.Delete(dbMap)
	if(err != nil) {
		return c.RenderJson(error.Error{Id:1,ErrorMessage:fmt.Sprint(err),DisplayMessage:"Some error occured"})
	}
	result := make(map[string]interface{})
	result["response"] = "ok"
	result["count"] = count
	result["note_id"] = note.Note_id
	return c.RenderJson(result)
}

func (c Note) checkErrors()(interface{}) {
	if(c.Validation.HasErrors()) {
		errors := make(map[string]interface{})
		errors["response"] = "error"
		error_count := len(c.Validation.Errors)
		errorMessages := make([]string,error_count,error_count)
		i := 0;
		for _,error := range c.Validation.Errors{
			errorMessages[i] = " " + error.Key + ":" + error.Message
			i++ 
		}
		errors["error"] = errorMessages
		errors["error_message"] = "Some fileds missing"
		return errors
	}
	return nil
}

func (c Note) GetUi() revel.Result {
	dbMap := database.NewConnection()
	defer dbMap.Db.Close()
	notes, _, err := model.GetNotes(dbMap)
	if(err != nil) {
		c.Validation.Error("Error occured while retreving data from db.")
	}
	return c.Render(notes)
}

func (c Note) AddUi() revel.Result {
	dbMap := database.NewConnection()
	defer dbMap.Db.Close()
	notes, _, err := model.GetNotes(dbMap)
	if(err != nil) {
		c.Validation.Error("Error occured while retreving data from db.")
	}
	return c.Render(notes)
}