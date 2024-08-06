const {
  connectDB,
  getObjectId,
  insertDocument,
  findDocument,
  findManyDocuments,
  updateDocument,
  deleteDocument,
  deleteManyDocuments,
} = require('../fs_driver/fs_driver');

let db;
let dbName = 'TaskApp';
//mongodb://tele-task-mongodb:27017/TaskApp
//mongodb://localhost:27017
async function initializeDB() {
  const uri = 'mongodb://tele-task-mongodb:27017'; // MongoDB connection string for localhost
  db = await connectDB(uri, dbName);
}

function getDB() {
  if (!db) throw new Error('Database not initialized');
  return db;
}

async function fs_create(collection, document) {
  return await insertDocument(db, collection, document);
}

async function fs_read(collection, query) {
  return await findDocument(db, collection, query);
}

async function fs_read_many(collection, query) {
  return await findManyDocuments(db, collection, query);
}

async function fs_update(collection, query, update) {
  return await updateDocument(db, collection, query, update);
}

async function fs_delete(collection, query) {
  return await deleteDocument(db, collection, query);
}

async function fs_deleteMany(collection, query) {
  return await deleteManyDocuments(db, collection, query);
}

module.exports = {
  initializeDB,
  getDB,
  fs_create,
  fs_read,
  fs_read_many,
  fs_update,
  fs_delete,
  fs_deleteMany,
  getObjectId,
};
