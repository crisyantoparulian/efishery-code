const db = require("../models");
const User = db.user;
const Role = db.role;
const Op = db.Sequelize.Op;

exports.allAccess = async (req, res) => {
    
    res.status(200).send({
        code: 200,
        message: "success",
        data: req.jwt_claim
    });
    return
};

exports.adminBoard = (req, res) => {
    res.status(200).send("Admin Content.");
  };