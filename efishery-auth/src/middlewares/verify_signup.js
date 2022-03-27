const db = require("../models");
const ROLES = db.ROLES;
const User = db.user;

checkDuplicateUserName = async (req, res, next) => {
    // Username
    let user = await User.findOne({
      where: {
        name: req.body.name
      }
    })

    if (user) {
      res.status(400).send({
        success: false,
        errorCode: "400",
        errorMessage: "Failed! username is already in use!",
        data: null
      });
      return;
    }

    user = await User.findOne({
      where: {
        phone: req.body.phone
      }
    })

    if (user) {
      res.status(400).send({
        success: false,
        errorCode: "400",
        errorMessage: "Failed! phone number is already in use!",
        data: null
      });
      return;
    }

    next();

  };

  const verifySignUp = {
    checkDuplicateUserName: checkDuplicateUserName,
  };
  module.exports = verifySignUp;