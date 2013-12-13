package model
import (
    "github.com/coopernurse/gorp"	
    "time"
	// "errors"
	// "fmt"
)

type Note struct {
	Note_id int `db:"note_id"`
	Content string	`db:"note"`
	Title string	`db:"title"`
	Created int64 `db:"created"`
	Modified int64 `db:"modified"`
}


func (n *Note) Save(dbMap *gorp.DbMap) error {
	n.Created = time.Now().Unix()
	n.Modified = time.Now().Unix()
	err := dbMap.Insert(n)
	if(err != nil) {
		return err
	}
	return nil
	//return errors.New("emit macho dwarf: elf header corrupted")
}

func (n *Note) Delete(dbMap *gorp.DbMap) (int64,error) {
	count,err := dbMap.Delete(n)
	if(err != nil) {
		return 0,err
	}
	return count,nil
	//return errors.New("emit macho dwarf: elf header corrupted")
}

func (n *Note) Update(dbMap *gorp.DbMap) (int64,error) {
	n.Modified = time.Now().Unix()
	count, err := dbMap.Update(n)
	if(err != nil) {
		return 0,err
	}
	return count,nil
	//return errors.New("emit macho dwarf: elf header corrupted")
}
func GetNotes(dbMap *gorp.DbMap) ([]Note, int, error) {
	//var notes []Note 
	notes := []Note{}
	_,err := dbMap.Select(&notes,"SELECT * FROM note")
	if err != nil {
		return nil,0,err
	}
	return notes,len(notes), nil

}