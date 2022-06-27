package gen

import (
	"strconv"

	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

var snowflakeNode *snowflake.Node

func Init() error {
	n, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}
	snowflakeNode = n
	return nil
}

// NewID use the simple method to create a new id
// TODO: add custom configuration for this section
func NewID() (id int64) {
	id = int64(snowflakeNode.Generate())
	zap.L().Info(strconv.Itoa(int(id)))
	return id
}

// func main() {
// 	n, err := snowflake.NewNode(1)
// 	if err != nil {
// 		println(err)
// 		os.Exit(1)
// 	}
//
// 	for i := 0; i < 3; i++ {
// 		id := n.Generate()
// 		fmt.Println("id", id)
// 		fmt.Println(
// 			"node: ", id.Node(),
// 			"step: ", id.Step(),
// 			"time: ", id.Time(),
// 			"\n",
// 		)
// 	}
// }
