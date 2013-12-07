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
func (c Note) Add() revel.Result {	
	dbMap := database.NewConnection()
    defer dbMap.Db.Close()
	note := model.Note{}
    c.Params.Bind(&note.Title,"title")
    c.Params.Bind(&note.Content,"content")
    fmt.Println(c.Params)
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