package messaging

import (
	"context"
	"fmt"

	"github.com/shotarosasaki/publisher/config"
	"github.com/shotarosasaki/publisher/domain"
	"google.golang.org/api/option"

	"cloud.google.com/go/pubsub"
)

// MEMO メッセージキューの性質を持つサービスへのメッセージ送信ロジックは、ここで吸収

type Publisher struct {
	client *pubsub.Client
	cfg    *config.QueueConfig
}

func NewPublisher(cfg *config.QueueConfig) (domain.Publisher, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, cfg.ProjectID, option.WithCredentialsFile(cfg.CredentialsPath))
	if err != nil {
		return nil, fmt.Errorf("Could not create pubsub Client; error: %v", err)
	}
	// TODO client.Close() をコールする必要性の確認と実行タイミング、コネクション保持方式か都度接続方式か要確認！

	pub := &Publisher{
		client: client,
		cfg:    cfg,
	}
	return pub, nil
}

func (p *Publisher) Publish(in *domain.PublishInput) (*domain.PublishOutput, error) {
	topic := p.client.Topic(p.cfg.TopicName)
	if topic == nil {
		return nil, fmt.Errorf("Could not get topic; topic: %s", p.cfg.TopicName)
	}
	ctx := context.Background()
	result := topic.Publish(ctx, &pubsub.Message{
		Data:       in.Data,
		Attributes: in.Attributes,
	})
	serverID, err := result.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("Could not get publish Result; error: %v", err)
	}

	return &domain.PublishOutput{ServerID: serverID}, nil
}
