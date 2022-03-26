const { body } = require("express-validator")

const validate = (method) => {
  switch (method) {
      case 'register': {
        return [
            body('name').not().isEmpty().withMessage("name must not empty."),
            body('role').not().isEmpty().withMessage("role must not empty."),
            body('phone').matches(/^(\+62|62|0)8[1-9][0-9]{6,9}$/).withMessage("phone number is not valid"),
          ]
    }
    case 'login': {
      return [
        body('password').not().isEmpty().isLength({min:4,max:4}),
        body('phone').matches(/^(\+62|62|0)8[1-9][0-9]{6,9}$/).withMessage("phone number is not valid"),
      ];
   }

  }
}

module.exports = validate;