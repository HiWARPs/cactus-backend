const asyncHandler = require('express-async-handler')
const Electron = require('../models/electrons')

const downloadResults = asyncHandler(async (req, res) => {
    let d = Date.now().valueOf()
    let fileName = d + "_results.csv";

    const electrons = await Electron.findOne()

    // TODO: convert to CSV content
    const fileContent = JSON.stringify(electrons.raw);
    const fileSize = Buffer.byteLength(fileContent, 'utf8');

    res.set({
            "Content-Disposition": `attachment; filename="${fileName}"`,
            "Content-Type": "text/plain",
            "Content-Length": fileSize
    });

    res.send(fileContent)
})

module.exports = {
    downloadFile: downloadResults,
}
