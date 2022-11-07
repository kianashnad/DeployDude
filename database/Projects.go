package database

import (
	"DeployDude/ent"
	"DeployDude/ent/project"
	"DeployDude/foundation"
	"context"
	"fmt"
	"log"
	"time"
)

func RemoveProject(hashID string) error {
	client := foundation.GetDBClient()
	defer client.Close()

	_, err := client.Project.
		Delete().
		Where(project.HashID(hashID)).
		Exec(context.Background())

	// if there is an error or there is no project with the hash_id provided
	if err != nil {
		return fmt.Errorf("failed getting project: %w", err)
	}

	return nil
}
func GetProjectByGitURL(gitURL string) (*ent.Project, error) {
	client := foundation.GetDBClient()
	defer client.Close()

	p, err := client.Project.
		Query().
		Where(project.GitURL(gitURL)).
		All(context.Background())

	// if there is an error or there is no project with the gitURL provided
	if err != nil || len(p) == 0 {
		return nil, fmt.Errorf("failed getting project: %w", err)
	}

	return p[0], nil
}

func GetProject(hashID string) (*ent.Project, error) {
	client := foundation.GetDBClient()
	defer client.Close()

	p, err := client.Project.
		Query().
		Where(project.HashID(hashID)).
		All(context.Background())

	// if there is an error or there is no project with the hash_id provided
	if err != nil || len(p) == 0 {
		return nil, fmt.Errorf("failed getting project: %w", err)
	}

	return p[0], nil
}

func CreateProject(GitURL string, DirPath string, Title string) (*ent.Project, error) {
	// hash_id: md5(timestamp, ProjectName,GitURL)
	projectHashID := foundation.GetMD5Hash(time.Now().String() + Title + GitURL)
	client := foundation.GetDBClient()
	defer client.Close()

	p, err := client.Project.
		Create().
		SetHashID(projectHashID).
		SetGitURL(GitURL).
		SetDirPath(DirPath).
		SetTitle(Title).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating project: %w", err)
	}
	log.Println("project was created: ", p)
	return p, nil
}
