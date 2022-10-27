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
app.get("/projects", getProjects);
app.get("/form/12", getForm);
app.post("/form/12/execute", executeForm);
app.get("/data", getData);
app.get("/test", getTestData)

app.listen(3001)
console.log("Node.js Express server is listening on port 3001...")


/////////////////////////////////
// Implementation of routes
//

function welcome(req, res) {
  res.json({ message: "☃️️ Welcome to Project Cactus. Our mock backend is ready for you.️ ☃️" });
}

function getProjects(req, res) {
  let filename = "mock_responses/projects.json";

  let rawData = fs.readFileSync(filename);
  let jsonData = JSON.parse(rawData);
  res.json(jsonData);
}

function getForm(req, res) {
  let filename = "mock_responses/form_12.json";

  let rawData = fs.readFileSync(filename);
  let jsonData = JSON.parse(rawData);
  res.json(jsonData);
}

function executeForm(req, res) {
  let filename = "mock_responses/form_12_query_data.json";
      
  let rawData = fs.readFileSync(filename);
  let jsonData = JSON.parse(rawData);
  res.json(jsonData);
}

function getData(req, res) {
  let filename = "mock_responses/data.json";
      
  let rawData = fs.readFileSync(filename);
  let jsonData = JSON.parse(rawData);
  res.json(jsonData);    
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
