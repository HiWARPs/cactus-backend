"use strict"

const fs = require('fs');
const express = require("express");
const app = express();
const mongojs = require('mongojs');
const collections = ['Electrons'];
const db = mongojs('test', collections);

// A security risk: Enabling CORS for development. But do NOT enable it for production.
const cors = require('cors');
app.use(cors());

/////////////////////////////////
// Routes
//
app.get("/", welcome);
app.get("/test", getTestData)

app.listen(3001)
console.log("Node.js Express server is listening on port 3001...")


/////////////////////////////////
// Implementation of routes
//

function welcome(req, res) {
  res.json({ message: "☃️️ Welcome to Project Cactus. Our mock backend is ready for you.️ ☃️" });
}

function getTestData(req, res) {
  const spin = 1;
  const min = 20;
  const max = 120;
  const inc = 20;
  db.Electrons.find({
    '$expr': {
      '$and': [
        {'$eq': ['$Spin', spin ]}, 
        {'$gte':['$Angle', min]},
        {'$lte':['$Angle', max]},
        {'$eq': [
          { '$mod': [{'$subtract': ['$Angle', min]}, inc]}, 0]
        },
      ]}
    },
    function(err, docs) {
      if (err) {
        res.send(err)
      } else {
        res.json(docs)
      }
    }
  )
}
