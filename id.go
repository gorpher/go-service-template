package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
)

type snowflakeID interface {
	Int64() int64
	String64() string
	String() string
	Generate() snowflake.ID
}

type id struct {
	node *snowflake.Node
}

func (i *id) Int64() int64 {
	return i.node.Generate().Int64()
}

func (i *id) String64() string {
	return fmt.Sprint(i.node.Generate().Int64())
}

func (i *id) String() string {
	return i.node.Generate().String()
}

func (i *id) Generate() snowflake.ID {
	return i.node.Generate()
}

var ID snowflakeID

func NewID() error {
	snowflake.Epoch = time.Now().Unix()
	rand.Seed(rand.Int63n(time.Now().UnixNano())) // nolint
	node := 0 + rand.Int63n(1023-0)               // nolint
	n, err := snowflake.NewNode(node)
	if err != nil {
		return err
	}
	ID = &id{n}
	return err
}
