const asyncHandler = require('express-async-handler')
const Project = require('../models/projects')

const getProjects = asyncHandler(async (req, res) => {
    const projects = await Project.find()

    res.status(200).json({message: projects})
})

const createProject = asyncHandler(async (req, res) => {

    const project = await Project.create({
        name: req.body.name,
    })

    res.status(200).json(project)
})

const deleteProject = asyncHandler(async (req, res) => {
    res.status(200).json({message: `DELETE not implemented, id: ${req.params.id}`})
})

const updateProject = asyncHandler(async (req, res) => {
    res.status(200).json({message: `UPDATE not implemented: id ${req.params.id}`})
})

module.exports = {
    getProjects,
    createProject,
    updateProject,
    deleteProject
}