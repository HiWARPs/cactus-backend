const asyncHandler = require('express-async-handler')
const Project = require("../models/projects");


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

const getForm = asyncHandler(async (req, res) => {
    formID = req.params.id

    const project = await Project.findById(req.params.pid)
    if (!project) {
        res.status(400)
        throw new Error("project not found")
    }

    const form = await project.forms.id(req.params.id)

    res.status(200).json(form)
})

module.exports = {
    getForm,
    createForm
}