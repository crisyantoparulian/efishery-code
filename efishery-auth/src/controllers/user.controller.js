const db = require("../models");
const User = db.user;
const Role = db.role;
const Op = db.Sequelize.Op;

exports.allAccess = async (req, res) => {
    
    res.status(200).send({
        success: true,
        errorCode: "",
        errorMessage: "success",
        data: req.jwt_claim
    });
    return
};