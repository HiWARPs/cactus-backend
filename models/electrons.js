const mongoose = require('mongoose');

const dataPointSchema = new mongoose.Schema({
    X1: { type: Number, required: true },
    X2: { type: Number, required: true },
    Y1: { type: Number, required: true }
});

const electronSchema = new mongoose.Schema({
    raw: { type: Object, required: true }, // How the client sent us the json. Useful for debugging our parser.
    dataPoints: { type: [dataPointSchema], required: true}
});

module.exports = mongoose.model('Electron', electronSchema)
