// +build !debug

package ffmpeg_go

func DebugNodes(node []DagNode) {}
// d
func DebugOutGoingMap(node []DagNode, m map[int]map[Label][]NodeInfo) {}
