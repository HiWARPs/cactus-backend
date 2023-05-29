const asyncHandler = require('express-async-handler')
const Electron = require('../models/electrons')
const ObjectsToCsv = require('objects-to-csv');

const downloadResults = asyncHandler(async (req, res) => {
    let d = Date.now().valueOf()
    let fileName = d + "_results.csv";

    const results = await Electron.findOne()

    if (results == null) {
        throw new Error('Results not found')
    }

    const csv = new ObjectsToCsv(results.raw);
    const fileContent = await csv.toString();
    const fileSize = Buffer.byteLength(fileContent, 'utf8');

    res.set({
        "Content-Disposition": `attachment; filename="${fileName}"`,
        "Content-Type": "text/plain",
        "Content-Length": fileSize
    });

    res.send(fileContent)
})

module.exports = {
    downloadResults: downloadResults,
}
