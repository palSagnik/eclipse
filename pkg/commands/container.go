package commands

type Container struct {
	ID      string
	ImageID string
	Created int64
	State   string
	Status  string
}
