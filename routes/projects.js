const express = require('express')
const router = express.Router();
const {
    getProjects,
    createProject,
    updateProject,
    deleteProject
} = require('../controllers/projects')

const {
    getForm,
} = require('../controllers/form')


router.route("/").get(getProjects).post(createProject)
router.route("/:id").put(updateProject).delete(deleteProject)

router.route("/:pid/form/:id").get(getForm)


module.exports = router;