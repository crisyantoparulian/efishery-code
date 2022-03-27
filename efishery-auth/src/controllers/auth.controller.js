const db = require("../models");
const config = require("../../config.json");
const randomString = require('../utils/random_string');
const User = db.user;
const Role = db.role;
const Op = db.Sequelize.Op;
var jwt = require("jsonwebtoken");
var bcrypt = require("bcryptjs");

exports.register = async (req, res) => {
    try {
        let pwd = randomString(4);
        // Save User to Database
        let user = await User.create({
            name: req.body.name,
            phone: req.body.phone,
            password: bcrypt.hashSync(pwd, 8)
        })
            
        let roles = await Role.findAll({
            where: {
            name: req.body.role
            }
        });
        if (roles.length == 0){
          let role = await Role.create({
            name: req.body.role,
          });
          roles = [role];
        }
        
        await user.setRoles(roles);
        
        return res.status(200).send({
            success: true,
            errorCode: "",
            errorMessage: "",
            data: {
              "id": user.id,
              "name": user.name,
              "phone": user.phone,
              "password": pwd,
              "role": req.body.role,
              "updated_at": user.updatedAt,
              "created_at": user.createdAt
            }
        });

    }catch (err) {
        console.log("Error Signup => ", err);
        return res.status(400).send({
            success: false,
            errorCode: "400",
            errorMessage: "Something wrong",
            data: null
        });
    }
};

exports.login = (req, res) => {
  User.findOne({
    where: {
      phone: req.body.phone
    }
  })
    .then(user => {

      if (!user) {
        return res.status(404).send({
            success: false,
            errorCode: "404",
            errorMessage: "User Not found.",
            data : null
        });
      }

      let passwordIsValid = bcrypt.compareSync(
        req.body.password,
        user.password
      );

      if (!passwordIsValid) {
        return res.status(401).send({
          success: false,
          errorCode: "401",
          errorMessage: "Invalid Password!",
          data : null
        });
      }

      user.getRoles().then(roles => {

        let claim = { 
            name: user.name,
            phone: user.phone,
            role: (roles.length > 0) ? roles[0].name : '',
            created_at: user.createdAt,
            updated_at: user.updatedAt
          }
        var token = jwt.sign(claim, config.jwt.secret, {
          expiresIn: 86400 // 24 hours
        });

        res.status(200).send({
          success: true,
          errorCode: "",
          errorMessage: "success",
            data: {
                id: user.id,
                name: user.name,
                phone: user.phone,
                role: claim.role,
                accessToken: token
            }
        });
      });

    })
    .catch(err => {
      res.status(500).send({
          success: false,
          errorCode: "500",
          errorMessage: "Invalid Password!",
          data : null
      });
    });
};