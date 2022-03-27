const { validationResult } = require('express-validator');

validateRequestJson = (req, res, next) => {
    // Finds the validation errors in this request and wraps them in an object with handy functions
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(400).json({
          success: false,
          errorCode: "400",
          errorMessage: "ValidationError",
          errors: errors.array(),
          data: null
        });
    }
    next()
  };

  const verifyRequestJson = {
    validateRequestJson: validateRequestJson,
  };
  module.exports = verifyRequestJson;