const { verifySignUp,verifyRequestJson,validate } = require("../middlewares");
// const { body, validationResult } = require('express-validator');
// const validate = require('../../middlewares/validate');
const controller = require("../controllers/auth.controller");

module.exports = function(app) {
  app.use(function(req, res, next) {
    res.header(
      "Access-Control-Allow-Headers",
      "x-access-token, Origin, Content-Type, Accept"
    );
    next();
  });

  app.post(
    "/api/v1/auth/register",
    [   
        validate('register'),
        validateRequestJson,
        verifySignUp.checkDuplicateUserName,
    ],
    controller.register
  );

  app.post("/api/v1/auth/login",validate('login'), validateRequestJson, controller.login);

};