package controller

import (
	"DeployDude/database"
	"DeployDude/foundation"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var password = foundation.GetENV("DEBU_MASTERPASSWORD")

func DeployProject(c *gin.Context) {
	var requestCtx foundation.DeployProjectDTO
	if err := c.ShouldBindJSON(&requestCtx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := database.GetProject(requestCtx.HashID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "Server Error 01",
		})
		return
	}

	foundation.DeployViaDockerCompose(project.DirPath, project.Title)
	defer foundation.PruneDocker()

	c.String(http.StatusOK, "Done!")
	return
}

func AddProject(c *gin.Context) {
	var requestCtx foundation.AddProjectDTO
	if err := c.ShouldBindJSON(&requestCtx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if requestCtx.Password != password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not remove the project. Code 10"})
		return
	}

	// Check if the project URL is from the {env DEDU_GITHUB_USERNAME} GitHub account:
	githubUsername := foundation.GetENV("DEDU_GITHUB_USERNAME")
	githubBaseURL := "git@github.com:"
	isGithubURL := requestCtx.GitURL[0:len(githubBaseURL)] == githubBaseURL
	isGithubUsernameValid := requestCtx.GitURL[len(githubBaseURL):len(githubBaseURL)+len(githubUsername)] == githubUsername

	if !isGithubURL || !isGithubUsernameValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Bad request.",
		})
		return
	}

	//Check if it doesn't exist:
	project, _ := database.GetProjectByGitURL(requestCtx.GitURL)
	if project != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Project already exists.",
		})
		return
	}

	// 1. Clone, 2. Generate hashID, 3. Return the Project hashID
	cmd := "cd " + foundation.GetENV("DEDU_BASE_DIR") + " && git clone " + requestCtx.GitURL
	_, err := foundation.RunCmd(cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "server error code 40",
		})
		return
	}

	folderName := strings.Split(strings.Split(requestCtx.GitURL, "/")[1], ".git")[0]
	createProject, err := database.CreateProject(requestCtx.GitURL, "/"+folderName+"/", requestCtx.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"error": "server error code 50",
		})
		return
	}
	defer foundation.PruneDocker()

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Project has been created successfully.",
		"hash_id": createProject.HashID,
	})
	return

}

func RemoveProject(c *gin.Context) {
	var requestCtx foundation.RemoveProjectDTO

	if err := c.ShouldBindJSON(&requestCtx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request."})
		return
	}

	if requestCtx.Password != password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not remove the project. Code 10"})
		return
	}

	project, err := database.GetProject(requestCtx.HashID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Could not remove the project. Code 50",
		})
		return
	}

	// removing the app from docker (with its volumes)
	err = foundation.RemoveAppFromDocker(project.DirPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Could not remove the project. Code 21",
		})
		return
	}

	//removing from the disk:
	_, err = foundation.RunCmd("cd " + foundation.GetENV("DEDU_BASE_DIR") + "rm -rf " + project.DirPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Could not remove the project. Code 20",
		})
		return
	}

	// after removing files, removing from the database:
	if err := database.RemoveProject(requestCtx.HashID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Could not remove the project. Code 60",
		})
		return
	}

	defer foundation.PruneDocker()

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Project has been removed successfully.",
	})
	return

}
