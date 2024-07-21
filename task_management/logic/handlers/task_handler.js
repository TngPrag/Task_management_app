const { taskCreateSchema} = require('../dto/task_create_dto');
const {taskUpdateStatusSchema, taskUpdateScheduleSchema} = require('../dto/task_update_dto')
const { TaskModel } = require('../core/task_model');
const { authenticate } = require('../../middlewares/middleware');
const { v4: uuidv4 } = require('uuid');
const { getUserRole } = require('../pkg/authz_proxy_role');
const { verifyPolicy } = require('../pkg/authz_proxy_policy');

/**
 * Create a new task.
 * @param {Object} req - The request object.
 * @param {Object} res - The response object.
 */
async function Task_write(req, res) {
  try {
    const { error, value } = taskCreateSchema.validate(req.body);

    if (error) {
      return res.status(400).json({ error: error.details[0].message });
    }

    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      return res.status(401).json({ error: 'Authorization header missing or invalid' });
    }

    const token = authHeader.split(' ')[1];
    const userProfile = await authenticate(token);

    const userRole = await getUserRole(token);
    if (userRole !== 'admin') {
      return res.status(403).json({ error: 'Unauthorized' });
    }

    const hasPolicy = await verifyPolicy(token, userProfile.user_id, 'task_app/task_manager_service/api/v0.1/task', 'POST');
    if (!hasPolicy) {
      return res.status(403).json({ error: 'Forbidden' });
    }

    const task_id = uuidv4();

    const taskData = {
      ...value,
      user_id: userProfile.user_id,
      owner_id: userProfile.user_id,
    };
    // notify the user about the task assignment
    try {
      await notifyUser(value.email, value.title, value.description, value.deadline, token);
    } catch (notifyError) {
      return res.status(500).json({ error: notifyError.message });
    }s
    const newTask = new TaskModel(taskData);
    const createdTask = await TaskModel.createTask(newTask);
    res.status(200).json(createdTask);
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
}

/**
 * Get a task by ID.
 * @param {Object} req - The request object.
 * @param {Object} res - The response object.
 */
async function Task_read(req, res) {
  try {
    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      return res.status(401).json({ error: 'Authorization header missing or invalid' });
    }

    const token = authHeader.split(' ')[1];
    const userProfile = await authenticate(token);

    const userRole = await getUserRole(token);
    if (userRole !== 'admin' && userRole !== 'user') {
      return res.status(403).json({ error: 'Unauthorized' });
    }

    const hasPolicy = await verifyPolicy(token, userProfile.user_id, 'task_app/task_manager_service/api/v0.1/task', 'GET');
    if (!hasPolicy) {
      return res.status(403).json({ error: 'Forbidden' });
    }

    const task = await TaskModel.getTask(req.params.id);
    if (!task) {
      return res.status(404).json({ error: 'Task not found' });
    }

    res.status(200).json(task);
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
}

/**
 * Update a task by ID.
 * @param {Object} req - The request object.
 * @param {Object} res - The response object.
 */
async function Task_update_schedule(req, res) {
  try {
    const { error, value } = taskUpdateScheduleSchema.validate(req.body);
    if (error) {
      return res.status(400).json({ error: error.details[0].message });
    }

    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      return res.status(401).json({ error: 'Authorization header missing or invalid' });
    }

    const token = authHeader.split(' ')[1];
    const userProfile = await authenticate(token);

    const userRole = await getUserRole(token);
    if (userRole !== 'admin') {
      return res.status(403).json({ error: 'Unauthorized' });
    }

    const hasPolicy = await verifyPolicy(token, userProfile.user_id, 'task_app/task_manager_service/api/v0.1/task', 'PUT');
    if (!hasPolicy) {
      return res.status(403).json({ error: 'Forbidden' });
    }

    const updatedTask = await TaskModel.updateTask(req.params.id, value);
    if (!updatedTask) {
      return res.status(404).json({ error: 'Task not found' });
    }

    res.status(200).json(updatedTask);
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
}
async function Task_update_status(req,res){
try {
    const { error, value } = taskUpdateStatusSchema.validate(req.body);
    if (error) {
      return res.status(400).json({ error: error.details[0].message });
    }

    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      return res.status(401).json({ error: 'Authorization header missing or invalid' });
    }

    const token = authHeader.split(' ')[1];
    const userProfile = await authenticate(token);

    const userRole = await getUserRole(token);
    if (userRole !== 'admin') {
      return res.status(403).json({ error: 'Unauthorized' });
    }

    const hasPolicy = await verifyPolicy(token, userProfile.user_id, 'task_app/task_manager_service/api/v0.1/task', 'PUT');
    if (!hasPolicy) {
      return res.status(403).json({ error: 'Forbidden' });
    }

    const updatedTask = await TaskModel.updateTask(req.params.id, value);
    if (!updatedTask) {
      return res.status(404).json({ error: 'Task not found' });
    }

    res.status(200).json(updatedTask);
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
}

/**
 * Delete a task by ID.
 * @param {Object} req - The request object.
 * @param {Object} res - The response object.
 */
async function Task_delete(req, res) {
  try {
    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      return res.status(401).json({ error: 'Authorization header missing or invalid' });
    }

    const token = authHeader.split(' ')[1];
    const userProfile = await authenticate(token);

    const userRole = await getUserRole(token);
    if (userRole !== 'admin') {
      return res.status(403).json({ error: 'Unauthorized' });
    }

    const hasPolicy = await verifyPolicy(token, userProfile.user_id, 'task_app/task_manager_service/api/v0.1/task', 'DELETE');
    if (!hasPolicy) {
      return res.status(403).json({ error: 'Forbidden' });
    }

    const deletedTask = await TaskModel.deleteTask(req.params.id);
    if (!deletedTask) {
      return res.status(404).json({ error: 'Task not found' });
    }

    res.status(200).json({ message: 'Task deleted successfully' });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
}

/**
 * List tasks by user ID.
 * @param {Object} req - The request object.
 * @param {Object} res - The response object.
 */
async function Task_listByUser(req, res) {
  try {
    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      return res.status(401).json({ error: 'Authorization header missing or invalid' });
    }

    const token = authHeader.split(' ')[1];
    const userProfile = await authenticate(token);

    const userRole = await getUserRole(token);
    if (userRole !== 'user') {
      return res.status(403).json({ error: 'Unauthorized' });
    }

    const hasPolicy = await verifyPolicy(token, userProfile.user_id, 'task_app/api/v0.1/task', 'GET');
    if (!hasPolicy) {
      return res.status(403).json({ error: 'Forbidden' });
    }

    const tasks = await TaskModel.listTasksByUserID(userProfile.user_id);
    res.status(200).json(tasks);
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
}

/**
 * List tasks by owner ID.
 * @param {Object} req - The request object.
 * @param {Object} res - The response object.
 */
async function Task_listByOwner(req, res) {
  try {
    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      return res.status(401).json({ error: 'Authorization header missing or invalid' });
    }

    const token = authHeader.split(' ')[1];
    const userProfile = await authenticate(token);

    const userRole = await getUserRole(token);
    if (userRole !== 'admin') {
      return res.status(403).json({ error: 'Unauthorized' });
    }

    const hasPolicy = await verifyPolicy(token, userProfile.user_id, 'task_app/task_manager_service/api/v0.1/task', 'GET');
    if (!hasPolicy) {
      return res.status(403).json({ error: 'Forbidden' });
    }

    const tasks = await TaskModel.listTasksByOwnerId(userProfile.user_id);
    res.status(200).json(tasks);
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
}

/**
 * Delete tasks by user ID.
 * @param {Object} req - The request object.
 * @param {Object} res - The response object.
 */
async function Task_deleteByUserId(req, res) {
  try {
    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      return res.status(401).json({ error: 'Authorization header missing or invalid' });
    }

    const token = authHeader.split(' ')[1];
    const userProfile = await authenticate(token);

    const userRole = await getUserRole(token);
    if (userRole !== 'admin') {
      return res.status(403).json({ error: 'Unauthorized' });
    }

    const hasPolicy = await verifyPolicy(token, userProfile.user_id, 'task_app/task_manager_service/api/v0.1/task', 'DELETE');
    if (!hasPolicy) {
      return res.status(403).json({ error: 'Forbidden' });
    }

    const deletedTasks = await TaskModel.deleteTaskByUserID(req.params.user_id);
    if (!deletedTasks) {
      return res.status(404).json({ error: 'Tasks not found' });
    }

    res.status(200).json({ message: 'Tasks deleted successfully' });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
}

/**
 * Delete tasks by owner ID.
 * @param {Object} req - The request object.
 * @param {Object} res - The response object.
 */
async function Task_deleteByOwnerId(req, res) {
  try {
    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      return res.status(401).json({ error: 'Authorization header missing or invalid' });
    }

    const token = authHeader.split(' ')[1];
    const userProfile = await authenticate(token);

    const userRole = await getUserRole(token);
    if (userRole !== 'admin') {
      return res.status(403).json({ error: 'Unauthorized' });
    }

    const hasPolicy = await verifyPolicy(token, userProfile.user_id, 'task_app/task_manager_service/api/v0.1/task', 'DELETE');
    if (!hasPolicy) {
      return res.status(403).json({ error: 'Forbidden' });
    }

    const deletedTasks = await TaskModel.deleteTaskByOwnerID(req.params.owner_id);
    if (!deletedTasks) {
      return res.status(404).json({ error: 'Tasks not found' });
    }

    res.status(200).json({ message: 'Tasks deleted successfully' });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
}

module.exports = { Task_write, Task_read, Task_update_schedule,Task_update_status, Task_delete, Task_listByUser,Task_listByOwner, Task_deleteByUserId, Task_deleteByOwnerId };
