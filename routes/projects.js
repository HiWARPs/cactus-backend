const express = require('express')
const router = express.Router();
const {
    getProjects,
    createProject,
    updateProject,
    deleteProject,
    getProjectByID
} = require('../controllers/projects')


router.route("/").get(getProjects).post(createProject)
router.route("/:id").get(getProjectByID).put(updateProject).delete(deleteProject)

module.exports = router;