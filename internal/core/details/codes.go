package details

const (
	HTTPMethodNotFound = 101
	HTTPContentInvalid = 102

	JSONInvalid      = 201
	ValidationFailed = 202
	DecoderFaulty    = 203

	RegistrationFailure    = 301
	RegistrationSuccessful = 302
	AuthFailure            = 303
	AuthSuccessful         = 304
	JWTTokenFailure        = 305
	JWTTokenSuccessful     = 306

	DocumentInsertionFailure    = 601
	DocumentInsertionSuccessful = 602
	IndexInsertionFailure       = 603
	IndexInsertionSuccessful    = 604
	DocumentFindFailure         = 605
	DocumentFindSuccessful      = 606
	NoDocumentFound             = 607

	EnvVariableNotDefined = 401

	MongoDBConnectionFailure         = 501
	MongoDBConnectionSuccessful      = 502
	MongoDBConnectionCloseFailure    = 503
	MongoDBConnectionCloseSuccessful = 504
	// MongoDBSchemaRegistrationSuccessful = 505
	// MongoDBSchemaRegistrationFailure    = 506
	MongoDBDocumentInsertionFailure    = 505
	MongoDBDocumentInsertionSuccessful = 506
	MongoDBIndexCreationFailure        = 507
	MongoDBIndexCreationSuccessful     = 508
	MongoDBFindSuccessful              = 509
	MongoDBFindEmpty                   = 510
)

var message = map[int]string{
	HTTPMethodNotFound: "Method Not Permitted",
	HTTPContentInvalid: "Content-Type/Accept is of invalid type",

	JSONInvalid:      "JSON is Invalid",
	ValidationFailed: "Validation Failed",
	DecoderFaulty:    "Decoder structure provided was faulty",

	RegistrationFailure:    "Registration failed",
	RegistrationSuccessful: "Registration successful",
	AuthFailure:            "Authentication failed",
	AuthSuccessful:         "Authentication successful",
	JWTTokenFailure:        "JWT Token creation failure",
	JWTTokenSuccessful:     "JWT Token create successful",

	DocumentInsertionFailure:    "Failed to insert document.",
	DocumentInsertionSuccessful: "Document inserted successfully.",
	IndexInsertionFailure:       "Failed to insert Index",
	IndexInsertionSuccessful:    "Index inserted successfully",
	DocumentFindFailure:         "Failed to find Document",
	DocumentFindSuccessful:      "Document found successfully",
	NoDocumentFound:             "No Document Found",

	EnvVariableNotDefined: "Environment Variable not defined",

	MongoDBConnectionFailure:           "MongoDB connection failed",
	MongoDBConnectionSuccessful:        "MongoDB connection successful",
	MongoDBConnectionCloseFailure:      "MongoDB connection termination failed",
	MongoDBConnectionCloseSuccessful:   "MongoDB connection termination Successfully",
	MongoDBDocumentInsertionFailure:    "MongoDB document failed during insertion",
	MongoDBDocumentInsertionSuccessful: "MongoDB document inserted successfully",
	MongoDBIndexCreationFailure:        "MongoDB Index creation failed.",
	MongoDBIndexCreationSuccessful:     "MongoDB Index created successfully",
	MongoDBFindSuccessful:              "MongoDB found document",
	MongoDBFindEmpty:                   "No MongoDB document found",
	// MongoDBSchemaRegistrationSuccessful: "Schema addition to MongoDB Successfully",
	// MongoDBSchemaRegistrationFailure:    "Schema addition to MongoDB Failed",

}

func GetMessage(code int) string {
	detail, found := message[code]
	if !found {
		return "Unknown Error Code"
	}
	return detail
}
