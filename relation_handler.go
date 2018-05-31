package main

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
)

// Relation between ID and OtherID, state can be liked | disliked | matched
type Relation struct {
	// pk means primary key
	ID      ID     `json:"-" sql:",pk"`
	OtherID ID     `json:"user_id" sql:",pk"`
	State   string `json:"state"`
	Type    string `json:"type" sql:"-"`
}

// RelationIndexHandler return one's all relations
func RelationIndexHandler(w http.ResponseWriter, r *http.Request) {
	defer serverErr(w)
	vars := mux.Vars(r)
	id := vars["user_id"]
	relations := make([]Relation, 0)
	err := userdb.Model(&relations).Where("id = ?", id).Select()
	checkErr(err)
	for i := range relations {
		relations[i].Type = "relationships"
	}
	json.NewEncoder(w).Encode(&relations)
}

// RelationUpdateHandler update relation between two users
func RelationUpdateHandler(w http.ResponseWriter, r *http.Request) {
	defer serverErr(w)
	decoder := json.NewDecoder(r.Body)
	var relation Relation
	err := decoder.Decode(&relation)
	checkErr(err)
	vars := mux.Vars(r)
	id64, err := strconv.ParseUint(vars["user_id"], 10, 64)
	checkErr(err)
	otherID64, err := strconv.ParseUint(vars["other_id"], 10, 64)
	checkErr(err)
	id := ID(id64)
	otherID := ID(otherID64)
	checkErr(err)
	err = userdb.RunInTransaction(updateState(id, otherID, relation.State))
	checkErr(err)
	relation, _ = selectRelation(id, otherID)
	relation.Type = "relationship"
	json.NewEncoder(w).Encode(&relation)
}

// db transaction function
func updateState(id, other ID, state string) func(*pg.Tx) error {
	return func(tx *pg.Tx) error {
		err := checkUserExist(id)
		if err != nil {
			return err
		}
		err = checkUserExist(other)
		if err != nil {
			return err
		}
		otherRelation, _ := selectRelation(other, id)
		newRelation := Relation{ID: id, OtherID: other, State: state}
		userdb.Model(&otherRelation).Where("id = ?", other).Where("other = ?", id).
			Select()
		switch {
		// if matched ignore. not support regret
		case otherRelation.State == "matched":
			return nil
		case otherRelation.State == "liked" && state == "liked":
			newRelation.State = "matched"
			otherRelation.State = "matched"
			_, err = userdb.Model(&otherRelation).Column("state").WherePK().Update()
			if err != nil {
				return err
			}
		default:
			// do nothing
		}
		err = userdb.Insert(&newRelation)
		// insert or update
		_, err = userdb.Model(&newRelation).OnConflict("(id, other_id) DO UPDATE").
			Set("state = EXCLUDED.state").Insert()
		return err
	}
}

func selectRelation(id, other ID) (Relation, error) {
	relation := Relation{}
	err := userdb.Model(&relation).Where("id = ?", id).Where("other_id = ?",
		other).Select()
	return relation, err
}
