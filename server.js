// create an express app
const express = require("express")
const app = express()
const mongojs = require('mongojs')

// use the express-static middleware
app.use(express.static("public"))

// define the first route
app.get("/", function (req, res) {
  res.send("<h1>Hello World!</h1>")
})


app.post("/projects", createProject)


// define the first route
app.get("/projects", function (req, res) {
  res.send("<h1>Hello World!</h1>")
})



// Connect to DB
const db = mongojs(process.env.DATABASE_URL, ['electrons'])

db.on('error', function (err) {
  console.log('database error', err)
})

db.on('connect', function () {
  console.log('database connected')
})


// start the server listening for requests
app.listen(process.env.PORT || 3000,
    () => console.log("Server is running..."));

function createProject(req, res) {
  const j = 5;
}
