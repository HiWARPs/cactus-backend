const mongoose = require('mongoose');

const dataPointSchema = new mongoose.Schema({
    X: { type: Number, required: true },
    YPi: { type: Number, required: true },
    YerrPi: { type: Number, required: true }
});

const electronSchema = new mongoose.Schema({
    raw: { type: Object, required: true }, // How the client sent us the json. Useful for debugging our parser.
    dataPoints: { type: [dataPointSchema], required: true}
});

module.exports = mongoose.model('Electron', electronSchema)
