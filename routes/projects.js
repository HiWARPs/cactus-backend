 const express = require('express')

 const router = express.Router();

 // define the first route
 router.get("/", function (req, res) {
  res.send("<h1>/projects</h1>")
 })



 module.exports = router;