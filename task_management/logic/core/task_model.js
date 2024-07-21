const {
  fs_create,
  fs_read,
  fs_find,
  fs_update,
  fs_delete,
  fs_deleteMany,
  fs_read_many,
} = require('../../fs/fs_conn');

// Define the schema for a task
class TaskModel {
  constructor(taskData) {
    //this.task_id = taskData.task_id;
    this.user_id = taskData.user_id;
    this.owner_id = taskData.owner_id;
    this.title = taskData.title;
    this.description = taskData.description;
    this.status = taskData.status;
    this.deadline = taskData.deadline;
  }
}

async function createTask(taskData) {
  return await fs_create('tasks', taskData);
}

async function getTask(id) {
  return await fs_read('tasks', { _id: id });
}

async function listTasksByUserId(userId) {
  return await fs_read_many('tasks', { user_id: userId });
}


async function updateTask(id, updateData) {
  return await fs_update('tasks', { _id: id }, updateData);
}

async function deleteTask(id) {
  return await fs_delete('tasks', { _id: id });
}

async function deleteTasksByUserId(userId) {
  return await fs_deleteMany('tasks', { user_id: userId });
}

async function deleteTasksByOwnerId(ownerId) {
  return await fs_deleteMany('tasks', { owner_id: ownerId });
}
async function listTasksByOwnerId(ownerId) {
  return await fs_read_many('tasks', { owner_id: ownerId });
}

module.exports = {
  createTask,
  getTask,
  listTasksByUserId,
  listTasksByOwnerId,
  updateTask,
  deleteTask,
  deleteTasksByUserId,
  deleteTasksByOwnerId,
};
