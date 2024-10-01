package domain

type Role int

const (
	ORGANIZATION_ADMIN int = iota
	PROJECT_LEADER
	COLLABORATOR
)
