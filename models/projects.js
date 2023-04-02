const mongoose = require('mongoose')

const xColumnsSchema = new mongoose.Schema( {
    name: {
      type: String,
    },
    range: {
      type: Boolean,
    },
})

const yColumnsSchema = new mongoose.Schema( {
    name: {
      type: String,
    },
})

const formSchema = new mongoose.Schema({
    name: {
        type: String,
    },
    description: {
      type: String,
    },
    references: {
      type: String,
    },
    xColumns: [xColumnsSchema],
    yColumns: [yColumnsSchema],
})

const projectsSchema = new mongoose.Schema({
    name: {
        type: String,
    },
    forms: [formSchema],
})

module.exports = mongoose.model('Project', projectsSchema)
