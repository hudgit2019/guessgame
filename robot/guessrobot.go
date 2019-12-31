package robot

import (
	"github.com/hudgit2019/leafboot/base"
	"github.com/hudgit2019/leafboot/robotlogic"
)

/*AI逻辑示例*/
type GuessRobotLogic struct {
	robotlogic.RobotLogic
}
type GuessRobotNode struct {
	robotlogic.RobotNode
	handCard []int8
}

func (g *GuessRobotLogic) OnCreateRobot() base.IPlayerNode {
	return &GuessRobotNode{}
}
func (g *GuessRobotLogic) RegisteRobotMsg() {
	g.RobotLogic.RegisteRobotMsg()
	g.RobotLogic.MapReqHandler["SendCard"] = g.handleSendCard
	g.RobotLogic.MapReqHandler["GameResult"] = g.handleGameResult
}
func (g *GuessRobotLogic) handleSendCard(userID int64, msg interface{}) {

}
func (g *GuessRobotLogic) handleGameResult(userID int64, msg interface{}) {

}
