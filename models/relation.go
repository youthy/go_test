/*Package models define custom structs and operations */
package models

import "github.com/go-pg/pg"

// Relation between ID and OtherID, state can be liked | disliked | matched
type Relation struct {
	// pk means primary key
	ID      ID     `json:"-" sql:",pk"`
	OtherID ID     `json:"user_id" sql:",pk"`
	State   string `json:"state"`
	Type    string `json:"type" sql:"-"`
}

// UpdateRelationState is db transaction function
func UpdateRelationState(id, other ID, state string) error {
	return Userdb.RunInTransaction(func(tx *pg.Tx) error {
		err := CheckUserExist(id)
		if err != nil {
			return err
		}
		err = CheckUserExist(other)
		if err != nil {
			return err
		}
		otherRelation, _ := GetRelation(other, id)
		newRelation := Relation{ID: id, OtherID: other, State: state}
		Userdb.Model(&otherRelation).Where("id = ?", other).Where("other = ?", id).
			Select()
		switch {
		// if matched ignore. not support regret
		case otherRelation.State == "matched":
			return nil
		case otherRelation.State == "liked" && state == "liked":
			newRelation.State = "matched"
			otherRelation.State = "matched"
			_, err = Userdb.Model(&otherRelation).Column("state").WherePK().Update()
			if err != nil {
				return err
			}
		default:
			// do nothing
		}
		err = Userdb.Insert(&newRelation)
		// insert or update
		_, err = Userdb.Model(&newRelation).OnConflict("(id, other_id) DO UPDATE").
			Set("state = EXCLUDED.state").Insert()
		return err
	})
}

// GetRelation return relation between id and other
func GetRelation(id, other ID) (Relation, error) {
	relation := Relation{}
	err := Userdb.Model(&relation).Where("id = ?", id).Where("other_id = ?",
		other).Select()
	return relation, err
}

// GetRelations select all relation of id's
func GetRelations(id ID) ([]Relation, error) {
	relations := make([]Relation, 0)
	err := Userdb.Model(&relations).Where("id = ?", id).Select()
	return relations, err
}
