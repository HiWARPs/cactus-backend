const express = require('express')
const router = express.Router();
const {
    getProjects,
    createProject,
    updateProject,
    deleteProject,
    getProjectByID
} = require('../controllers/projects')

const {
    getForm,
    createForm,
    updateForm,
    deleteForm,
} = require('../controllers/form')


router.route("/").get(getProjects).post(createProject)
router.route("/:id").get(getProjectByID).put(updateProject).delete(deleteProject)

router.route("/:pid/form").post(createForm)
router.route("/:pid/form/:id").get(getForm).put(updateForm).delete(deleteForm)

module.exports = router;
