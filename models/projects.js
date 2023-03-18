const mongoose = require('mongoose')

const formSchema = new mongoose.Schema({
    name: {
        type: String,
    }
})

const projectsSchema = new mongoose.Schema({
    name: {
        type: String,
    },
    forms: [formSchema],
})


module.exports = mongoose.model('Project', projectsSchema)