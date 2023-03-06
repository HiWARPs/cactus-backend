const mongoose = require('mongoose')

const projectsSchema = new mongoose.Schema({
    name: {
        type: String,
    }
})

module.exports = mongoose.model('Project', projectsSchema)