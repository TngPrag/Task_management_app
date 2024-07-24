/**
 * @swagger
 * /health:
 *   get:
 *     summary: Check the health status of the service
 *     tags: [Health]
 *     responses:
 *       200:
 *         description: Service is running
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 status:
 *                   type: string
 *                   example: 'Task manager service is running Ok!'
 *       500:
 *         description: Internal server error
 */
async function checkHealth(req, res) {
    res.json({ status: 'Task manager service is running Ok!' });
  }
  
  module.exports = {
    checkHealth,
  };
  