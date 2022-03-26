const express = require("express");
const cors = require("cors");
const app = express();
// var db = require("./database.js");
const db = require("./src/models");
const config = require("./src/config/config.js");

// enable cors
app.use(cors());
app.options('*', cors());
// parse requests of content-type - application/json

app.use(express.json());
// parse requests of content-type - application/x-www-form-urlencoded

app.use(express.urlencoded({ extended: true }));
// simple route

app.get("/", (req, res) => {
  res.json({ message: "Welcome to efishery jwt apps" });
});


const Role = db.role;
db.sequelize.sync({force: true}).then(() => {
    console.log('Drop and Resync Db');
    initial();
  });

  // routes
require('./src/routes/auth.routes')(app);
require('./src/routes/user.routes')(app);


app.use(function(req, res, next) {
	return res.status(404).send({
        code:404,
        message : "NotFound"
    });
});

// set port, listen for requests
const PORT = config.app.port || 8080;
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}.`);
});

function initial() {   
    Role.create({
      id: 1,
      name: "admin"
    });
  }