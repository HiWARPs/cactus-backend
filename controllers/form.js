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

const updateForm = asyncHandler(async (req, res) => {

  const projectId = req.params.pid
  const formId = req.params.id

  const project = await Project.findById(projectId)
  if (!project) {
    res.status(400)
    throw new Error("project not found")
  }

  const form = await project.forms.id(formId)
  if (!form) {
    res.status(400)
    throw new Error("form not found")
  }

  form.name = req.body.name;
  form.description = req.body.description;
  form.references = req.body.references;
  form.xColumns = req.body.xColumns;
  form.yColumns = req.body.yColumns;
  const updatedForm = await project.save();

  res.status(200).json(updatedForm)
})

module.exports = {
    getForm,
    createForm,
    updateForm,
}
