const mongoose = require('mongoose')

const referencesSchema = new mongoose.Schema( {
  name: {
    type: String,
    required: true,
  },
})

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
    references: [referencesSchema],
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
