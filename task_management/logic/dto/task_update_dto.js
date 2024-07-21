const Joi = require('joi');

const taskUpdateStatusSchema = Joi.object({
  status: Joi.string().valid('pending', 'in-progress', 'completed').required(),
});

const taskUpdateScheduleSchema = Joi.object({
    deadline: Joi.date().required()
});

module.exports = {
  taskUpdateScheduleSchema,
  taskUpdateStatusSchema,
};
