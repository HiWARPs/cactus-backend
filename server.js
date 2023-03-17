// create an express app
const express = require('express')
const app = express()
const port = process.env.PORT || 3000
const colors = require('colors')


const connectDB = require('./db')
connectDB()

app.use(express.static("public"))
app.use(express.json())
app.use(express.urlencoded({extended: false}))

// define the first route
app.get("/", function (req, res) {
    res.status(200).json({"message": "top directory"})
})

// Import routes
const projectRoutes = require('./routes/projects')
app.use('/project', projectRoutes)


// Connect to DB
// mongoose.connect(process.env.DATABASE_URL);

// start the server listening for requests
app.listen(port,
    () => console.log("Server is running..."));