# API Documentation

## 📌 **Register Endpoint**

**Route:** `POST /register`

**Description:** Registers a new tenant with the server.

---

### 📨 Request Body

```json
{
  "org_name": "your_organization_name",
  "email": "contact_email@example.com",
  "service": "desired_service_name"
}
```

| Field      | Type   | Description                              |
| ---------- | ------ | ---------------------------------------- |
| `org_name` | String | The name of your organization            |
| `email`    | String | A valid email address for registration   |
| `service`  | String | The name of the service to be registered |

---

### ❓ Example Request

```http
POST /register HTTP/1.1
Content-Type: application/json

{
    "org_name" : "test",
    "email": "test1@test.com",
    "service": "testing"
}
```

### ✅ Example Response

```json
{
    "HTTPStatusCode": 200,
    "message": "Registation Completed",
    "details": {
        "Error": false,
        "statusCode": 302,
        "message": "Registration successful",
        "additionalDetails": {
            "api_key": "<API KEY>"
        }
    }
}
```

---

## 📌 **Auth Endpoint**

**Route:** `POST /auth/token`

**Description:** Authenticates tenant with the server and provide them with a JWT Token.

---

### 📨 Request Body

```json
{
    "api_key": "<API KEY>"
}
```

| Field      | Type   | Description                              |
| ---------- | ------ | ---------------------------------------- |
| `api_key`  | String | API Key recieved from registration       |

---

### ❓ Example Request

```http
POST /auth/token HTTP/1.1
Content-Type: application/json

{
    "api_key": "<API KEY>"
}
```

### ✅ Example Response

```json
{
    "HTTPStatusCode": 200,
    "message": "Authentication Completed",
    "details": {
        "Error": false,
        "statusCode": 304,
        "message": "Authentication successful",
        "additionalDetails": {
            "token": "<JWT TOKEN>"
        }
    }
}
```

---

## 📌 **Schema Endpoint**

**Route:** `POST /schema`

**Description:** Registers a new tenant with the server.

---

### 📨 Request Body

```json
{
    "collections":[
        {
            "name": "name of collection",
            "properties": {
                "field1": {
                    "bsonType": "type for the field",
                    "description": "description for the field",
                    "minimum": "Minimum value",
                    "maximum": "Maximum value"
                },
            "required": ["field1"]
            }
        }
    ]
}
```
Here's your updated Markdown documentation using the `### 📨 Request Body` section style, including the example JSON and a detailed table for each field and subfield:

---

### 📨 Request Body

```json
{
  "collections": [
    {
      "name": "name of collection",
      "properties": {
        "field1": {
          "bsonType": "type for the field",
          "description": "description for the field",
          "minimum": "Minimum value",
          "maximum": "Maximum value"
        }
      },
      "required": ["field1"]
    }
  ]
}
```

---

### 📋 Field Descriptions

| Field         | Type | Description                                                                 | Required |
| ------------- | ---- | --------------------------------------------------------------------------- | -------- |
| `collections` | List | List of collections that the tenant will define. Must contain at least one. | ✅ Yes   |

---

#### 🗂️ Collection Object (`collections[]`)

| Field        | Type                   | Description                                                                                              | Required |
| ------------ | ---------------------- | -------------------------------------------------------------------------------------------------------- | -------- |
| `name`       | String                 | Name of the collection.                                                                                  | ✅ Yes   |
| `properties` | Map (`string: object`) | A map of field names to their definitions. Each key is the field name, and the value is a `fieldSchema`. | ✅ Yes   |
| `required`   | List of Strings        | List of field names that are required in documents of this collection.                                   | ❌ No    |

---

#### 🔧 Field Object (`properties[fieldName]`)

| Field         | Type    | Description                                               | Required     |
| ------------- | ------- | --------------------------------------------------------- | ------------ |
| `bsonType`    | String  | The BSON data type for the field (e.g., `string`, `int`). | ✅ Yes       |
| `description` | String  | A brief description of the field.                         | ❌ No        |
| `minimum`     | Integer | Minimum value (only for numeric types).                   | ❌ No        |
| `maximum`     | Integer | Maximum value (only for numeric types).                   | ❌ No        |


---

### ❓ Example Request

```http
POST /schema HTTP/1.1
Authorization: Bearer <JWT TOKEN>
Content-Type: application/json

{
    "collections": [
        {
            "name": "testingcollection",
            "properties": {
                "name": {
                    "bsonType": "string",
                    "description": "name"
                },
                "age": {
                    "bsonType": "int",
                    "minimum": 2,
                    "maximum": 10
                }
            },
            "required": [
                "name"
            ]
        }
    ]
}
```

### ✅ Example Response

```json
{
    "HTTPStatusCode": 200,
    "message": "Schema Registation Completed",
    "details": {
        "Error": false,
        "statusCode": 307,
        "message": "Schema Registered Successfully"
    }
}
```

---
