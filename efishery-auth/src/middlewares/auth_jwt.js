const jwt = require("jsonwebtoken");
const config = require("../config/config.js");
const db = require("../models");
const User = db.user;

verifyToken = (req, res, next) => {
    let token = req.headers["x-access-token"];
    if (!token) {
        return res.status(403).send({
            code : 403,
            message: "No token provided!",
        });
    }
    jwt.verify(token, config.jwt.secret, (err, decoded) => {
        if (err) {
        return res.status(401).send({
            code : 401,
            message: "Unauthorized!",
        });
        }
        req.jwt_claim = decoded;
        next();
    });
};

isAdmin = async (req, res, next) => {
    try {
        let user = await User.findOne( {where: {
            name: req.jwt_claim.name
          }})
        if (user){
            let roles = await user.getRoles();
            if (roles.length > 0){
                for (let i = 0; i < roles.length; i++) {
                    if (roles[i].name === "admin") {
                    next();
                    return;
                    }
                }
            }
            res.status(403).send({
                code:403,
                message: "require admin role!"
            });
        }else{
            res.status(403).send({
                code:403,
                message: "user not valid!"
            });
        }
    }catch(err){
        console.log("Error middleware is admin",err);
        res.status(500).send({
            code:500,
            message: "SomethinWrong"
        });
    }
    
};

const authJwt = {
  verifyToken: verifyToken,
  isAdmin: isAdmin,
};
module.exports = authJwt;