const jwt = require("jsonwebtoken");
const config = require("../../config.json");
const db = require("../models");
const User = db.user;

verifyToken = (req, res, next) => {
    let token = req.headers["Authorization"] || req.headers["authorization"];
    if (!token || token == '') {
        return res.status(403).send({
            success: false,
            errorCode: "403",
            errorMessage: "No token provided!",
            data: null
        });
    }

    token = token.replace(/^Bearer\s+/, "");
    
    if (!token) {
        return res.status(403).send({
            success: false,
            errorCode: "403",
            errorMessage: "No token provided!",
            data: null
        });
    }
    jwt.verify(token, config.jwt.secret, (err, decoded) => {
        if (err) {
        return res.status(401).send({
            success: false,
            errorCode: "401",
            errorMessage: "Unauthorized!",
            data: null
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
                success: false,
                errorCode: "403",
                errorMessage: "require admin role!",
                data: null
            });
        }else{
            res.status(403).send({
                success: false,
                errorCode: "403",
                errorMessage: "user not valid!",
                data: null
            });
        }
    }catch(err){
        console.log("Error middleware is admin",err);
        res.status(500).send({
            success: false,
            errorCode: "500",
            errorMessage: "SomethinWrong",
            data: null
        });
    }
    
};

const authJwt = {
  verifyToken: verifyToken,
  isAdmin: isAdmin,
};
module.exports = authJwt;