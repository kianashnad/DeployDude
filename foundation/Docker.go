package foundation

func PruneDocker() error {
	_, err := RunCmd("docker system prune -a --volumes")
	if err != nil {
		return err
	}
	return nil
}

func RemoveAppFromDocker(relativePath string) error {
	// stops containers, removes them, removes the volumes associated to them
	_, err := RunCmd("cd " + GetENV("DEDU_BASE_DIR") + relativePath + " && docker compose rm -f -s -v")
	if err != nil {
		return err
	}

	return nil
}
