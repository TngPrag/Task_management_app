const { taskCreateSchema } = require('../dto/task_create_dto');
const { taskUpdateStatusSchema, taskUpdateScheduleSchema } = require('../dto/task_update_dto');
const { 
  createTask, 
  getTask, 
  listTasksByUserId, 
  listTasksByOwnerId, 
  updateTask, 
  deleteTask, 
  deleteTasksByUserId, 
  deleteTasksByOwnerId 
} = require('../core/task_model');  // Ensure this path is correct
const { authenticate } = require('../../middlewares/middleware');
const { v4: uuidv4 } = require('uuid');
const { getUserRole } = require('../pkg/authz_proxy_role');
const { verifyPolicy } = require('../pkg/authz_proxy_policy');
const { notifyUser } = require('../pkg/user_notify');

/**
 * Create a new task.
 * @param {Object} req - The request object.
 * @param {Object} res - The response object.
 */
/**
 * @swagger
 * tags:
 *   name: Tasks
 *   description: Task management operations
 */

/**
 * @swagger
 * /task/write:
 *   post:
 *     summary: Create a new task
 *     tags: [Tasks]
 *     security:
 *       - BearerAuth: []
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               title:
 *                 type: string
 *               description:
 *                 type: string
 *               deadline:
 *                 type: string
 *     responses:
 *       200:
 *         description: Task created successfully
 *       400:
 *         description: Invalid input
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       500:
 *         description: Internal server error
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
    console.log(userRole)
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
      //user_id: userProfile.user_id,
      owner_id: userProfile.user_id,
    };
    
    // notify the user about the task assignment
    try {
      await notifyUser(value.user_id,value.title, value.description, value.deadline, token);
    } catch (notifyError) {
      return res.status(500).json({ error: notifyError.message });
    }
    //const newTask = new TaskModel(taskData);
    const createdTask = await createTask(taskData);
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

/**
 * @swagger
 * /task/read/{id}:
 *   get:
 *     summary: Get a task by ID
 *     tags: [Tasks]
 *     security:
 *       - BearerAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         description: Task ID
 *         schema:
 *           type: string
 *     responses:
 *       200:
 *         description: Task found
 *       404:
 *         description: Task not found
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       500:
 *         description: Internal server error
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

    const task = await getTask(req.params.id);
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

/**
 * @swagger
 * /task/update_schedule/{id}:
 *   put:
 *     summary: Update the schedule of a task
 *     tags: [Tasks]
 *     security:
 *       - BearerAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         description: Task ID
 *         schema:
 *           type: string
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               deadline:
 *                 type: string
 *                 format: date-time
 *     responses:
 *       200:
 *         description: Task schedule updated successfully
 *       400:
 *         description: Invalid input
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       404:
 *         description: Task not found
 *       500:
 *         description: Internal server error
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

    const updatedTask = await updateTask(req.params.id, value);
    if (!updatedTask) {
      return res.status(404).json({ error: 'Task not found' });
    }

    res.status(200).json(updatedTask);
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
}

/**
 * @swagger
 * /task/update_status/{id}:
 *   put:
 *     summary: Update the status of a task
 *     tags: [Tasks]
 *     security:
 *       - BearerAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         description: Task ID
 *         schema:
 *           type: string
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               status:
 *                 type: string
 *                 enum: [pending, in-progress, completed]
 *     responses:
 *       200:
 *         description: Task status updated successfully
 *       400:
 *         description: Invalid input
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       404:
 *         description: Task not found
 *       500:
 *         description: Internal server error
 */
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

    const updatedTask = await updateTask(req.params.id, value);
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

/**
 * @swagger
 * /task/remove/{id}:
 *   delete:
 *     summary: Delete a task by ID
 *     tags: [Tasks]
 *     security:
 *       - BearerAuth: []
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         description: Task ID
 *         schema:
 *           type: string
 *     responses:
 *       200:
 *         description: Task deleted successfully
 *       404:
 *         description: Task not found
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       500:
 *         description: Internal server error
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

    const deletedTask = await deleteTask(req.params.id);
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

/**
 * @swagger
 * /tasks/list_by_user:
 *   get:
 *     summary: List tasks assigned to the authenticated user
 *     tags: [Tasks]
 *     security:
 *       - BearerAuth: []
 *     responses:
 *       200:
 *         description: List of tasks
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       500:
 *         description: Internal server error
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

    const tasks = await listTasksByUserId(userProfile.user_id);
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

/**
 * @swagger
 * /tasks/list_by_admin:
 *   get:
 *     summary: List all tasks owned by the authenticated admin
 *     tags: [Tasks]
 *     security:
 *       - BearerAuth: []
 *     responses:
 *       200:
 *         description: List of tasks
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       500:
 *         description: Internal server error
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

    const tasks = await listTasksByOwnerId(userProfile.user_id);
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

/**
 * @swagger
 * /tasks/remove_by_user:
 *   delete:
 *     summary: Delete tasks assigned to a specific user
 *     tags: [Tasks]
 *     security:
 *       - BearerAuth: []
 *     parameters:
 *       - in: query
 *         name: user_id
 *         required: true
 *         description: User ID
 *         schema:
 *           type: string
 *     responses:
 *       200:
 *         description: Tasks deleted successfully
 *       404:
 *         description: Tasks not found
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       500:
 *         description: Internal server error
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

    const deletedTasks = await deleteTasksByUserId(req.params.user_id);
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
/**
 * @swagger
 * /tasks/remove_by_owner:
 *   delete:
 *     summary: Delete tasks owned by a specific owner
 *     tags: [Tasks]
 *     security:
 *       - BearerAuth: []
 *     parameters:
 *       - in: query
 *         name: owner_id
 *         required: true
 *         description: Owner ID
 *         schema:
 *           type: string
 *     responses:
 *       200:
 *         description: Tasks deleted successfully
 *       404:
 *         description: Tasks not found
 *       401:
 *         description: Unauthorized
 *       403:
 *         description: Forbidden
 *       500:
 *         description: Internal server error
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

    const deletedTasks = await deleteTasksByOwnerId(req.params.owner_id);
    if (!deletedTasks) {
      return res.status(404).json({ error: 'Tasks not found' });
    }

    res.status(200).json({ message: 'Tasks deleted successfully' });
  } catch (err) {
    res.status(500).json({ error: err.message });
  }
}

module.exports = { Task_write, Task_read, Task_update_schedule,Task_update_status, Task_delete, Task_listByUser,Task_listByOwner, Task_deleteByUserId, Task_deleteByOwnerId };
