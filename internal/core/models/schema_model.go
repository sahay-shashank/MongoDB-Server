package models

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/sahay-shashank/mongodb-server/internal/core/details"
	"github.com/sahay-shashank/mongodb-server/internal/core/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type DeleteSchemaRequest struct {
	Collection string `json:"collection"`
}

type schema struct {
	Collection []collectionSchema `json:"collections" validate:"required,min=1,dive"`
}

type collectionSchema struct {
	CollectionName string                 `json:"name" validate:"required"`
	Properties     map[string]fieldSchema `json:"properties" validate:"required,min=1,dive"`
	Required       []string               `json:"required,omitempty" validate:"omitempty,min=1,dive,required"`
}

type fieldSchema struct {
	BsonType    string `json:"bsonType" validate:"required"`
	Description string `json:"description,omitempty"`
	Minimum     *int   `json:"minimum,omitempty"`
	Maximum     *int   `json:"maximum,omitempty"`
}

type SchemaInterface interface {
	getCollectionList() []collectionSchema
	ConvertToValidator() (map[string]bson.M, details.APIDetails)
	GetCollectionNames() []string
}

func NewSchemaModel(data []byte) (SchemaInterface, details.APIDetails) {
	var schemas schema

	if err := json.Unmarshal(data, &schemas); err != nil {
		// 	return nil, fmt.Errorf("error unmarshalling JSON Data: %v", err)
		// }
		// validate := validator.New()
		// if err := validate.Struct(access); err != nil {
		// 	return nil, fmt.Errorf("missing fields: %v", err)
		// }
		return nil, details.APIDetails{
			Error:             true,
			StatusCode:        details.JSONInvalid,
			Message:           details.GetMessage(details.JSONInvalid),
			AdditionalDetails: err,
		}
	}
	if err := utils.Validate.Struct(schemas); err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Field '%s' failed validation: %s", err.Field(), err.Tag())
			errorMessages = append(errorMessages, errorMessage)
		}
		return nil, details.APIDetails{
			Error:             true,
			StatusCode:        details.ValidationFailed,
			Message:           details.GetMessage(details.ValidationFailed),
			AdditionalDetails: errorMessages,
		}
	}
	return &schemas, details.APIDetails{
		StatusCode: details.SchemaCreationSuccess,
		Message:    details.GetMessage(details.SchemaCreationSuccess),
	}
}

func (schema *schema) ConvertToValidator() (map[string]bson.M, details.APIDetails) {
	validator := make(map[string]bson.M, 0)
	for _, collectionSchema := range schema.getCollectionList() {
		processingResult := processSchema(&collectionSchema, validator)
		if processingResult.Error {
			return nil, details.APIDetails{
				Error:             true,
				StatusCode:        details.SchemaConversionFailure,
				Message:           details.GetMessage(details.SchemaConversionFailure),
				AdditionalDetails: processingResult,
			}
		}
	}
	return validator, details.APIDetails{
		StatusCode: details.SchemaConversionSuccess,
		Message:    details.GetMessage(details.SchemaConversionSuccess),
	}
}

func (a *schema) getCollectionList() []collectionSchema {
	return a.Collection
}

func (a *schema) GetCollectionNames() []string {
	collectionNames := make([]string, 0)
	for _, collection := range a.getCollectionList() {
		collectionNames = append(collectionNames, collection.getCollectionName())
	}
	return collectionNames
}

func processSchema(collectionSchema *collectionSchema, validator map[string]bson.M) details.APIDetails {
	fieldValidatorResult := collectionSchema.fieldValidator()
	if fieldValidatorResult.Error {
		return details.APIDetails{
			Error:             true,
			StatusCode:        details.ValidatorSchemaProccessingFailure,
			Message:           details.GetMessage(details.ValidatorSchemaProccessingFailure),
			AdditionalDetails: fieldValidatorResult,
		}
	}
	validator[collectionSchema.getCollectionName()] = collectionSchema.toValidator()

	return details.APIDetails{
		StatusCode: details.ValidatorSchemaProccessingSuccess,
		Message:    details.GetMessage(details.ValidatorSchemaProccessingSuccess),
	}
}

func (c *collectionSchema) getCollectionName() string {
	return c.CollectionName
}

func (c *collectionSchema) fieldValidator() details.APIDetails {
	for _, fields := range c.Required {
		if _, found := c.Properties[fields]; !found {
			return details.APIDetails{
				Error:             true,
				StatusCode:        details.SchemaFieldNotFound,
				Message:           details.GetMessage(details.SchemaFieldNotFound),
				AdditionalDetails: fmt.Errorf("%s field not found in properties", fields),
			}

		}
	}
	return details.APIDetails{
		StatusCode: details.SchemaFieldsFound,
		Message:    details.GetMessage(details.SchemaFieldsFound),
	}
}

func (c *collectionSchema) toValidator() bson.M {
	PropertiesBSONData := bson.M{}
	for fieldName, fieldSchema := range c.Properties {
		fieldBSON := bson.M{
			"bsonType": fieldSchema.BsonType,
		}

		if fieldSchema.Description != "" {
			fieldBSON["description"] = fieldSchema.Description
		}
		if fieldSchema.Minimum != nil {
			fieldBSON["minimum"] = *fieldSchema.Minimum
		}
		if fieldSchema.Maximum != nil {
			fieldBSON["maximum"] = *fieldSchema.Maximum
		}
		PropertiesBSONData[fieldName] = fieldBSON
	}
	validatorBSON := bson.M{
		"$jsonSchema": bson.M{
			"bsonType":   "object",
			"title":      fmt.Sprintf("%s Validator", c.CollectionName),
			"properties": PropertiesBSONData,
		},
	}
	if len(c.Required) > 0 {
		validatorBSON["$jsonSchema"].(bson.M)["required"] = c.Required
	}
	return validatorBSON
}
