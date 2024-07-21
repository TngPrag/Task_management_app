import { expect } from 'chai';
import mockRequire from 'mock-require';
import {
  createTask,
  getTask,
  listTasksByUserId,
  listTasksByOwnerId,
  updateTask,
  deleteTask,
  deleteTasksByUserId,
  deleteTasksByOwnerId
} from '../logic/core/task_model.js';

// Mock the fs module
const fsMock = {
  fs_create: async (collection, data) => data,
  fs_read: async (collection, query) => ({ _id: query._id, title: 'Test Task' }),
  fs_read_many: async (collection, query) => [
    { _id: 'task1', user_id: query.user_id || query.owner_id },
    { _id: 'task2', user_id: query.user_id || query.owner_id }
  ],
  fs_update: async (collection, query, updateData) => updateData,
  fs_delete: async (collection, query) => ({ deletedCount: 1 }),
  fs_deleteMany: async (collection, query) => ({ deletedCount: 2 })
};

// Mock the actual module
mockRequire('../fs/fs_conn.js', fsMock);

describe('TaskModel functions', () => {
  after(() => {
    mockRequire.stopAll();
  });

  describe('createTask', () => {
    it('should create a new task', async () => {
      const taskData = {
        user_id: 'user123',
        owner_id: 'owner123',
        title: 'Test Task',
        description: 'This is a test task',
        status: 'pending',
        deadline: '2024-07-30'
      };

      const result = await createTask(taskData);

      expect(result).to.deep.equal(taskData);
    });
  });

  describe('getTask', () => {
    it('should retrieve a task by id', async () => {
      const taskId = 'task123';
      const taskData = { _id: taskId, title: 'Test Task' };

      const result = await getTask(taskId);

      expect(result).to.deep.equal(taskData);
    });
  });

  describe('listTasksByUserId', () => {
    it('should list tasks by user id', async () => {
      const userId = 'user123';
      const tasks = [
        { _id: 'task1', user_id: userId },
        { _id: 'task2', user_id: userId }
      ];

      const result = await listTasksByUserId(userId);

      expect(result).to.deep.equal(tasks);
    });
  });

  describe('listTasksByOwnerId', () => {
    it('should list tasks by owner id', async () => {
      const ownerId = 'owner123';
      const tasks = [
        { _id: 'task1', owner_id: ownerId },
        { _id: 'task2', owner_id: ownerId }
      ];

      const result = await listTasksByOwnerId(ownerId);

      expect(result).to.deep.equal(tasks);
    });
  });

  describe('updateTask', () => {
    it('should update a task by id', async () => {
      const taskId = 'task123';
      const updateData = { title: 'Updated Task' };

      const result = await updateTask(taskId, updateData);

      expect(result).to.deep.equal(updateData);
    });
  });

  describe('deleteTask', () => {
    it('should delete a task by id', async () => {
      const taskId = 'task123';

      const result = await deleteTask(taskId);

      expect(result).to.deep.equal({ deletedCount: 1 });
    });
  });

  describe('deleteTasksByUserId', () => {
    it('should delete tasks by user id', async () => {
      const userId = 'user123';

      const result = await deleteTasksByUserId(userId);

      expect(result).to.deep.equal({ deletedCount: 2 });
    });
  });

  describe('deleteTasksByOwnerId', () => {
    it('should delete tasks by owner id', async () => {
      const ownerId = 'owner123';

      const result = await deleteTasksByOwnerId(ownerId);

      expect(result).to.deep.equal({ deletedCount: 2 });
    });
  });
});
