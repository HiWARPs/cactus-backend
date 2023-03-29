const express = require("express");
const router = express.Router();

const {
    uploadFile,
} = require('../controllers/files')

const multer  = require('multer');
let uploadFileMiddleware = multer({ dest: 'uploads/' }).single("file")

router.route("/").post(uploadFileMiddleware, uploadFile)


module.exports = router;
