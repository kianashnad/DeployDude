package foundation

type AddProjectDTO struct {
	Title    string `json:"title" binding:"required"`
	Password string `json:"password" binding:"required"`
	GitURL   string `json:"git_url" binding:"required"`
}

type RemoveProjectDTO struct {
	Password string `json:"password" binding:"required"`
	HashID   string `json:"hash_id" binding:"required"`
}

type DeployProjectDTO struct {
	HashID string `json:"hash_id" binding:"required"`
}
