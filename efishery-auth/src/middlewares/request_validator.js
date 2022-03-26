const { validationResult } = require('express-validator');

validateRequestJson = (req, res, next) => {
    // Finds the validation errors in this request and wraps them in an object with handy functions
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(400).json({
          code:400,
          message : "ValidationError",
          errors: errors.array()
        });
    }
    next()
  };

  const verifyRequestJson = {
    validateRequestJson: validateRequestJson,
  };
  module.exports = verifyRequestJson;