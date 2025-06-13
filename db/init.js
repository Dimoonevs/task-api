

const dbName = 'taskdb';
const coll    = 'tasks';

db = db.getSiblingDB(dbName);

db.createCollection(coll, {
    validator: {
        $jsonSchema: {
            bsonType: 'object',
            required: ['title', 'status', 'created_at', 'updated_at'],
            properties: {
                _id:         { bsonType: 'string' },
                title:       { bsonType: 'string',  description: 'short title' },
                description: { bsonType: 'string' },
                status: {
                    enum: [
                        'todo',
                        'in_progress',
                        'in_qa',
                        'ready_for_release',
                        'done'
                    ],
                    description: 'task lifecycle state'
                },
                created_at:  { bsonType: 'long' },
                updated_at:  { bsonType: 'long' }
            }
        }
    }
});

db[coll].createIndex({ status: 1 });
db[coll].createIndex({ created_at: -1 });
