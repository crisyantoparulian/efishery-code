const db = require("../models");
const config = require("../config/config.js");
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
        })
        await user.setRoles(roles);
        
        user.password = pwd;
        return res.status(200).send({
            code: 200,
            message: "User was registered successfully!",
            data: user
        });

    }catch (err) {
        console.log("Error Signup => ", err);
        return res.status(400).send({
            code:400,
            message:"SomethingWrong"
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
            message: "User Not found." 
        });
      }

      let passwordIsValid = bcrypt.compareSync(
        req.body.password,
        user.password
      );

      if (!passwordIsValid) {
        return res.status(401).send({
          code:401,
          message: "Invalid Password!",
          data : {
            accessToken: null,
          }
        });
      }

      user.getRoles().then(roles => {
        // let authorities = [];
        // for (let i = 0; i < roles.length; i++) {
        //   authorities.push("ROLE_" + roles[i].name.toUpperCase());
        // }

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
            code: 200,
            message: "success",
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
      res.status(500).send({ message: err.message });
    });
};