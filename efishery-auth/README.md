Efishery Authentication Apps
This apps is for authentication user

Config
You can copy file config.json.example into config.json for each apps

Application
This Apps consist of 2 parts

Authentication apps (mandatory)
Fetch apps (mandatory)

Run the application locally

For Auth apps, you can go inside efishery-auth folder and run below command

$ npm install

$ npm run start

For Fetch apps, you can go inside fetch-app folder and run below command

$ go run main.go fetchapp


$ go run main.go webapp
You may adjust the host path for mongodb location on your local

Run the application from docker
You need to rename the host into mongodb. From this

{
    "host": "localhost",
    "database": "paxelinclubdb"
}
into this
{
    "host": "mongodb",
    "database": "paxelinclubdb"
}
then run the

./command.sh
Whats Next?
Need to define the port of each apps
Define the username password for the mongodb
moving the config file into the external apps