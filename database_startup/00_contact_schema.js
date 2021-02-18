conn = Mongo();
db = conn.getDB("manager");

db.createCollection("contacts",{
    validator: {
        $jsonSchema: {
            bsonType: "object",
            required: ["name1"],
            properties: {
                name1: {
                    bsonType: "string",
                    description: "First name. This property is required"
                },
                name2: {
                    bsonType: "string",
                    description: "Second name"
                },
                surname: {
                    bsonType: "string",
                    description: "Surname"
                },
                email: {
                    bsonType: "string",
                    description: "Email address"
                },
                phone: {
                    bsonType: "string",
                    description: "Phone number"
                },
                website: {
                    bsonType: "string",
                    description: "Website url"
                },
                createdAt: {
                    bsonType: "string",
                    description: "Timestamp of creation date"
                },
                updatedAt: {
                    bsonType: "string",
                    description: "Timestamp of update date"
                }
            }
               
        }
    }
});




