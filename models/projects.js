const mongoose = require('mongoose')

const xColumnsSchema = new mongoose.Schema( {
    name: {
      type: String,
      required: true,
    },
    range: {
      type: Boolean,
    },
})

const yColumnsSchema = new mongoose.Schema( {
    name: {
      type: String,
      required: true,
    },
})

const formSchema = new mongoose.Schema({
    name: {
        type: String,
        required: true,
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
        required: true,
    },
    forms: [formSchema],
})

module.exports = mongoose.model('Project', projectsSchema)
