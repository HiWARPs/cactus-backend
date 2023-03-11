// create an express app
const express = require('express')
const app = express()
const port = process.env.PORT || 3000
const colors = require('colors')


const connectDB = require('./db')
connectDB()

// use the express-static middleware
app.use(express.static("public"))
app.use(express.json())
app.use(express.urlencoded({extended: false}))

// define the first route
app.get("/", function (req, res) {
    res.status(200).json({"message": "top directory"})
})

// Import routes
const projectRoutes = require('./routes/project')
app.use('/project', projectRoutes)


// start the server listening for requests
app.listen(port,
    () => console.log("Server is running..."));