const asyncHandler = require('express-async-handler')
const Project = require("../models/projects");

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
}