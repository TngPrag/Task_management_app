const { MongoClient, ObjectId } = require('mongodb');

async function connectDB(uri, dbName) {
  const client = new MongoClient(uri);
  await client.connect();
  return client.db(dbName);
}

function getObjectId(id) {
  return new ObjectId(id);
}

async function insertDocument(db, collection, document) {
  const result = await db.collection(collection).insertOne(document);
  return result;
}

async function findDocument(db, collection, query) {
  const result = await db.collection(collection).findOne(query);
  return result;
}

async function findManyDocuments(db, collection, query) {
  const result = await db.collection(collection).find(query).toArray();
  return result;
}

async function updateDocument(db, collection, query, update) {
  const result = await db.collection(collection).updateOne(query, { $set: update });
  return result;
}

async function deleteDocument(db, collection, query) {
  const result = await db.collection(collection).deleteOne(query);
  return result;
}

async function deleteManyDocuments(db, collection, query) {
  const result = await db.collection(collection).deleteMany(query);
  return result;
}

module.exports = {
  connectDB,
  getObjectId,
  insertDocument,
  findDocument,
  findManyDocuments,
  updateDocument,
  deleteDocument,
  deleteManyDocuments,
};
