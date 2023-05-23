const express = require("express");
const router = express.Router();

const {
    downloadResults,
} = require('../controllers/results')

router.route("/").get(downloadResults)
module.exports = router;
