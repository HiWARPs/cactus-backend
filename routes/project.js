const express = require('express')
const router = express.Router();
const {
    createProject,
    getProjects,
    getProjectByID,
    updateProject,
    deleteProject
} = require('../controllers/project')


router.route("/").get(getProjects).post(createProject)
router.route("/:id").get(getProjectByID).put(updateProject).delete(deleteProject)

module.exports = router;