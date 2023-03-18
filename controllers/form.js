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

module.exports = {
    getForm,
}