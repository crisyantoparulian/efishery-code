const db = require("../models");
const ROLES = db.ROLES;
const User = db.user;

checkDuplicateUserName = (req, res, next) => {
    // Username
    User.findOne({
      where: {
        name: req.body.name
      }
    }).then(user => {
      if (user) {
        res.status(400).send({
          code : 400,
          message: "Failed! Username is already in use!",
          data: null,
        });
        return;
      }
      next();
    });
  };

  const verifySignUp = {
    checkDuplicateUserName: checkDuplicateUserName,
  };
  module.exports = verifySignUp;