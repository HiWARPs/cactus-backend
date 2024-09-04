package controllers

import (
	"testing"

	"github.com/HiWARPs/cactus-backend/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFindForm(t *testing.T) {
	id := primitive.NewObjectID()
	hasIDList := []models.Form{
		{ID: id},
	}
	f, err := findForm(id, hasIDList)
	assert.Nil(t, err)
	assert.NotNil(t, f)
}

func TestChangeForms(t *testing.T) {
	assert := assert.New(t)

	id := primitive.NewObjectID()
	formsList := []models.Form{
		{ID: id, Name: "Name", Description: "Description", References: "References"},
		{ID: primitive.NewObjectID(), Name: "Name", Description: "Description", References: "References"},
	}

	changes := map[string]interface{}{
		"name": "New Name",
	}

	changedForms := changeForms(id, formsList, changes)
	changedForm, err := findForm(id, changedForms)

	assert.Nil(err)
	assert.Equal(changedForm.Name, "New Name")
}
