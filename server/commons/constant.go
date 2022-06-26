package commons

const (
	ConnectionTimeout   = 90 // seconds
	DatabaseName        = "hzp_task"
	MigrationFolderPath = "file://server/repositories/db/migrations"
	EndpointPath        = "/api/v1/taskmanagement"
)

type RepositoryType string

const (
	RepositoryTypeDB       RepositoryType = "DB"
	RepositoryTypeFile     RepositoryType = "File"
	RepositoryTypeInMemory RepositoryType = "InMemory"
)
