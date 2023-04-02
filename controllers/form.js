const asyncHandler = require('express-async-handler')
const Project = require("../models/projects");

const getForm = asyncHandler(async (req, res) => {
    const projectId = req.params.pid
    const formId = req.params.id

    const project = await Project.findById(projectId)
    if (!project) {
        res.status(400)
        throw new Error("project not found")
    }

    const form = await project.forms.id(formId)

    res.status(200).json(form)
})

const createForm = asyncHandler(async (req, res) => {
  const project = await Project.findById(req.params.pid)

  if (!project) {
    res.status(400)
    throw new Error("project not found")
  }

  project.forms.push(req.body)
  project.save();

  res.status(200).json(project)
})

module.exports = {
    getForm,
    createForm,
}
