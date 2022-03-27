module.exports = (sequelize, Sequelize) => {
    const Role = sequelize.define("roles", {
        id: {
            type: Sequelize.INTEGER,
            primaryKey: true,
            autoIncrement: true,
        },
        name: {
            type: Sequelize.STRING
        },
        createdAt: { type: Sequelize.DATE, field: 'created_at' },
        updatedAt: { type: Sequelize.DATE, field: 'updated_at' },
    });
    return Role;
  };