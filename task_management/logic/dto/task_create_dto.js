const Joi = require('joi');

const taskCreateSchema = Joi.object({
  user_id: Joi.string().required(),
  title: Joi.string().required(),
  description: Joi.string().required(),
  status: Joi.string().valid('pending', 'in-progress', 'completed').required(),
  deadline: Joi.date().required(),
});

module.exports = {
  taskCreateSchema,
};
