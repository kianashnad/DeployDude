package foundation

func DeployViaDockerCompose(projectDir string, projectName string) {
	SendAppLog(GetENV("DEDU_LOG_TEXT_NEW_DEPLOYMENT"), projectName)
	_, err := RunCmd("cd " + GetENV("DEDU_BASE_DIR") + projectDir + " && git pull origin main && docker compose up -d --build --force-recreate")
	if err != nil {
		SendAppLog(GetENV("DEDU_LOG_TEXT_NEW_DEPLOYMENT_FAILED"), projectName)
	} else {
		SendAppLog(GetENV("DEDU_LOG_TEXT_NEW_DEPLOYMENT_DONE"), projectName)
	}
}
