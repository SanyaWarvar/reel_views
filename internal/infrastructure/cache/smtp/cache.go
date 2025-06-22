package smtp

import (
	"context"
	"encoding/json"
	"rv/internal/infrastructure/cache/common"
	"rv/pkg/applogger"
	"rv/pkg/database/dragonfly"
	"rv/pkg/util"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	logger applogger.Logger
	client *dragonfly.Client
}

func NewCache(logger applogger.Logger, client *dragonfly.Client) *Cache {
	return &Cache{
		logger: logger,
		client: client,
	}
}

type ConfirmationCode struct {
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"createdAt"`
}

func (ch *Cache) SaveConfirmCode(ctx context.Context, email, code string, ttl *time.Duration) error {
	data, err := json.Marshal(ConfirmationCode{Code: code, CreatedAt: util.GetCurrentUTCTime()})
	if err != nil {
		return err
	}
	return ch.client.Save(ctx, common.EmailConfirmationCodes, email, data, ttl)
}

func (ch *Cache) GetConfirmCode(ctx context.Context, email string) (string, time.Duration, bool, error) {
	var ttl time.Duration
	var code string
	data, err := ch.client.GetOne(ctx, common.EmailConfirmationCodes, email)
	if err != nil {
		switch err {
		case redis.Nil:
			return code, ttl, false, nil

		}
		return code, ttl, false, err
	}

	var output ConfirmationCode
	err = json.Unmarshal(data, &output)
	if err != nil {
		return code, ttl, false, err
	}
	return output.Code, time.Now().UTC().Sub(output.CreatedAt), true, nil

}
