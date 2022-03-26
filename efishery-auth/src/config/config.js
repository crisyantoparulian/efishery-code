module.exports = {
    app : {
      port : 8081
    },
    database :{
      host: "0.0.0.0",
      user: "root",
      password: "",
      db: "database",
      dialect: "sqlite",
      pool: {
        max: 5,
        min: 0,
        acquire: 30000,
        idle: 10000
      },
      storage: "db.sqlite"
    },
    jwt : {
      secret : "cris-key"
    }
  };