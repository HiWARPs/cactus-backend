const asyncHandler = require('express-async-handler')
const path = require('path');

const downloadFile = asyncHandler(async (req, res) => {
    let file = "myfile.txt";
    let fileLocation = path.join('./uploads',file);
    console.log(fileLocation);
    res.download(fileLocation, file);
})

module.exports = {
    downloadFile,
}