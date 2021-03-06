const express = require("express");
const cors = require("cors");
const app = express();
const db = require("./src/models");
const config = require('./config.json');
const swaggerUi = require('swagger-ui-express'),
    swaggerDocument = require('./swagger.json');


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

db.sequelize.sync({force: false}).then(() => {
    console.log('Drop and Resync Db');
    // initial();
});

  // routes
require('./src/routes/auth.routes')(app);
require('./src/routes/user.routes')(app);

app.use('/swagger', swaggerUi.serve, function(req, res) {
    swaggerDocument.host = req.get('host'); // Replace hardcoded host information in Swagger file
    swaggerDocument.schemes = [req.protocol]; // Replace hardcoded protocol information in Swagger file
    swaggerUi.setup(swaggerDocument)(req, res);
});

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