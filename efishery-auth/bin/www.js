'use strict';

// SERVER DEFINITION

const chalk = require('chalk');
const http = require('http');
const app = require('../app').app;
const Sequelize = require('../models').sequelize;
const shipmentDB = require('../app').shipmentDB;
const memberDB = require('../app').memberDB;

const port = normalizePort(process.env.PORT || '3000');
const server = http.createServer(app);

// Set Server Timeout
server.timeout = 1000 * 60 * 5; // 5 minute

// Server listen

server.listen(port);
server.on('error', onError);
server.on('listening', onListening);

// HTTP event listener on 'Error'

function onError(err) {
	if (err.syscall !== 'listen') {
		throw err;
	}

	let bind = typeof port === 'string' ? `Pipe ${port}` : `Port ${port}`;

	switch (err.code) {
		case 'EACCES':
			console.error(`${bind} requires elevated privileges`);
			process.exit(1);
			break;
		case 'EADDRINUSE':
			console.error(`${bind} is already in use`);
			setTimeout(() => {
				server.close();
				process.exit(1);
			}, 2500);
			break;
		default:
			console.error(err);
			setTimeout(() => {
				server.close();
				process.exit(1);
			}, 2500);
	}
}

// HTTP event listener on 'Listening'

function onListening() {
	let addr = server.address();
	let bind = typeof addr === 'string' ? `Pipe ${addr}` : `port ${addr.port}`;
	console.log(' %s Server listening on %s in %s mode', chalk.green('✓'), bind, app.settings.env);
}

// Normalize port into number or string

function normalizePort(val) {
	let port = parseInt(val, 10);

	if (isNaN(port)) {
		return val;
	}
	if (port >= 0) {
		return port;
	}
	return false;
}

// Graceful Shutdown

process.on('SIGINT', () => {

	server.close(() => {
		console.log('Stop receiving new connection.. Should restart worker');
		shipmentDB.close(() => {
			console.log('Shipment DB closed');
			memberDB.close(() => {
				console.log('Member DB closed');
				Sequelize.close().then(() => {
					console.log('SQL DB closed');
					process.exit();
				});
			})
		});
	});

	setTimeout(function() {
		console.log('Timeout passed.. force shutdown');
		process.exit(1);
	}, 3500);
});