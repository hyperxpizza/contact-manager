conn = Mongo();
db = conn.getDB("manager");

db.createCollection("users", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      properties: {
        username: {
            bsonType: "string",
            description: "username"
        },
        passwordHash: {
            bsonType: "string",
            description: "Hashed string of users password"
        },
        email: {
            bsonType: "string",
            description: "User's email"
        },
        isAdmin: {
            bsonType: "bool",
            description: "Indication if user is an administrator or not"
        },
        createdAt: {
            bsonType: "string",
            description: "Timestamp of creation date"
        },
        updatedAt: {
            bsonType: "string",
            description: "Timestamp of update date"
        }
      },
    },
  },
});