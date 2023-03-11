const asyncHandler = require('express-async-handler')
const Project = require('../models/project')

const getProjects = asyncHandler(async (req, res) => {
    const projects = await Project.find()

    res.status(200).json({message: projects})
})

const getProjectByID = asyncHandler(async (req, res) => {
    const id = req.params.id;

    try {
        const item = await Project.findById(id);
        if (!item) {
            return res.status(404).send('Item not found');
        }
        res.json(item);
    } catch (error) {
        console.error(error);
        res.status(500).send('Server error');
    }
})


const createProject = asyncHandler(async (req, res) => {
    const project = await Project.create(req.body)
    res.status(200).json(project)
})

const deleteProject = asyncHandler(async (req, res) => {

    const itemId = req.params.id;

    try {
        const deletedItem = await Project.findByIdAndDelete(itemId);
        if (!deletedItem) {
            return res.status(404).send('Item not found');
        }
        res.send('Item deleted successfully');
    } catch (error) {
        console.error(error);
        res.status(500).send('Server error');
    }
})

const updateProject = asyncHandler(async (req, res) => {

    const project = await Project.findById(req.params.id)
    if (!project) {
        res.status(400)
        throw new Error("project not found")
    }

    const updatedProject = await Project.findByIdAndUpdate(req.params.id,
        req.body, {new: true})

    res.status(200).json(updatedProject)
})

module.exports = {
    getProjects,
    getProjectByID,
    createProject,
    updateProject,
    deleteProject
}