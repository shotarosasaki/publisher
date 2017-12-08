package domain

type Publisher interface {
	Publish(in *PublishInput) (*PublishOutput, error)
}

type PublishInput struct {
	Data       []byte
	Attributes map[string]string
}

type PublishOutput struct {
	ServerID string
}
