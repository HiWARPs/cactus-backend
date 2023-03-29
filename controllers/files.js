const asyncHandler = require('express-async-handler')
const csv = require("csvtojson");
const Electron = require('../models/electrons')

const uploadFile = asyncHandler(async (req, res) => {
    const jsonData=await csv().fromFile(req.file.path);
    const electrons = new Electron({ raw: jsonData, dataPoints: jsonData });

    electrons.save()
        .then(() => {
            res.status(201).json("successfully saved document")
        })
        .catch((error) => {
            res.status(500).json("failed to save document" + error)
        })
})

module.exports = {
    uploadFile,
}