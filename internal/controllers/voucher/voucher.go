package voucher

import (
	"time"

	"github.com/PauloLucas94/fase-4-hf-voucher/internal/entities/voucher"
	"github.com/PauloLucas94/fase-4-hf-voucher/internal/repository/adapter"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
)

type Controller struct {
	repository adapter.Interface
}

type Interface interface {
	ListOne(ID uuid.UUID) (entity voucher.Voucher, err error)
	ListAll() (entities []voucher.Voucher, err error)
	Create(entity *voucher.Voucher) (uuid.UUID, error)
	Update(ID uuid.UUID, entity *voucher.Voucher) error
	Remove(ID uuid.UUID) error
}

func NewController(repository adapter.Interface) Interface {
	return &Controller{repository: repository}
}

func (c *Controller) ListOne(id uuid.UUID) (entity voucher.Voucher, err error) {
	entity.ID = id
	response, err := c.repository.FindOne(entity.GetFilterId(), entity.TableName())
	if err != nil {
		return entity, err
	}
	return voucher.ParseDynamoAtributeToStruct(response.Item)
}

func (c *Controller) ListAll() (entities []voucher.Voucher, err error) {
	entities = []voucher.Voucher{}
	var entity voucher.Voucher

	filter := expression.Name("code").NotEqual(expression.Value(""))
	condition, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return entities, err
	}

	response, err := c.repository.FindAll(condition, entity.TableName())
	if err != nil {
		return entities, err
	}

	if response != nil {
		for _, value := range response.Items {
			entity, err := voucher.ParseDynamoAtributeToStruct(value)
			if err != nil {
				return entities, err
			}
			entities = append(entities, entity)
		}
	}

	return entities, nil
}

func (c *Controller) Create(entity *voucher.Voucher) (uuid.UUID, error) {
	entity.CreatedAt = time.Now()
	_, err := c.repository.CreateOrUpdate(entity.GetMap(), entity.TableName())
	return entity.ID, err
}

func (c *Controller) Update(id uuid.UUID, entity *voucher.Voucher) error {
	found, err := c.ListOne(id)
	if err != nil {
		return err
	}
	found.ID = id
	found.Code = entity.Code
	found.Percentage = entity.Percentage
	found.UpdatedAt = time.Now().AddDate(0, 1, 0)
	_, err = c.repository.CreateOrUpdate(found.GetMap(), entity.TableName())
	return err
}

func (c *Controller) Remove(id uuid.UUID) error {
	entity, err := c.ListOne(id)
	if err != nil {
		return err
	}
	_, err = c.repository.Delete(entity.GetFilterId(), entity.TableName())
	return err
}
