package tools

import (
	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
	"math/rand"
)

var (
	node *snowflake.Node
	err  error
)

func Init() error {

	// Create a new Node with a random Node number
	node, err = snowflake.NewNode(rand.Int63n(1024))
	if err != nil {
		logrus.Errorf("failed at creating new snowflake node, err: %v", err)
		return err
	}

	return nil
}

func GenSnowflakeID() int64 {
	if node == nil {
		node, err = snowflake.NewNode(rand.Int63n(1024))
	}

	// Generate a snowflake ID.
	id := node.Generate()

	return id.Int64()
}
