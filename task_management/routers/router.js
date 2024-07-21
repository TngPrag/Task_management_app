//const express = require('express');
const {checkHealth} = require('../logic/handlers/service_health')
const express = require('express');
const router = express.Router();

const {
  Task_write,
  Task_read,
  Task_delete,
  Task_listByOwner,
  Task_deleteByUserId,
  Task_deleteByOwnerId,
  Task_update_status,
  Task_update_schedule
} = require('../logic/handlers/task_handler');
//router.('/task_app/task_manager_service/api/v0.1/');

// Route for service health check
router.get('/health', checkHealth);
router.post('/task/write', Task_write);
router.get('/task/read/:id', Task_read);
router.put('/task/update_status/:id', Task_update_status);
router.put('/task/update_schedule/:id', Task_update_schedule);
router.delete('/task/remove/:id', Task_delete);
router.get('/tasks/list_by_admin', Task_listByOwner);
router.delete('/tasks/list_by_user', Task_deleteByUserId);
router.delete('/tasks/remove_by_owner', Task_deleteByOwnerId);
router.delete('/tasks/remove_by_user', Task_deleteByUserId)
module.exports = router;

