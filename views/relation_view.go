package views

import (
	"work/models"
)

// RelationIndex is view of relations index route
func RelationIndex(relations []models.Relation) []models.Relation {
	for i := range relations {
		relations[i] = RelationShow(relations[i])
	}
	return relations
}

// RelationShow is view of relation
func RelationShow(relation models.Relation) models.Relation {
	relation.Type = "relationships"
	return relation
}
