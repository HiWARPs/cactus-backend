const express = require("express");
const router = express.Router();

const {
    downloadFile,
} = require('../controllers/download')

router.route("/").get(downloadFile)
module.exports = router;
