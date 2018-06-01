package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"work/models"
	"work/util"
	"work/views"
)

// RelationIndexHandler return one's all relations
func RelationIndexHandler(w http.ResponseWriter, r *http.Request) {
	defer util.ServerErr(w)
	vars := mux.Vars(r)
	id64, err := strconv.ParseUint(vars["user_id"], 10, 32)
	if err != nil {
		util.BadRequest(w)
		return
	}
	relations, err := models.GetRelations(models.ID(id64))
	util.CheckErr(err)
	views.Render(w, views.RelationIndex(relations))
}

// RelationUpdateHandler update relation between two users
func RelationUpdateHandler(w http.ResponseWriter, r *http.Request) {
	defer util.ServerErr(w)
	decoder := json.NewDecoder(r.Body)
	var relation models.Relation
	err := decoder.Decode(&relation)
	util.CheckErr(err)
	vars := mux.Vars(r)
	id64, err := strconv.ParseUint(vars["user_id"], 10, 32)
	if err != nil {
		util.BadRequest(w)
		return
	}
	otherID64, err := strconv.ParseUint(vars["other_id"], 10, 32)
	if err != nil {
		util.BadRequest(w)
		return
	}
	id := models.ID(id64)
	otherID := models.ID(otherID64)
	err = models.UpdateRelationState(id, otherID, relation.State)
	util.CheckErr(err)
	relation, _ = models.GetRelation(id, otherID)
	views.Render(w, views.RelationShow(relation))
}
