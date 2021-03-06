package constants

const (
	RunCloudcore = "cd ${GOPATH}/src/github.com/kubeedge/kubeedge/cloud/; sudo nohup ./cloudcore > cloudcore.log 2>&1 &"
	RunEdgecore  = "cd ${GOPATH}/src/github.com/kubeedge/kubeedge/edge/; sudo nohup ./edgecore > edgecore.log 2>&1 &"
	RunEdgeSite  = "cd ${GOPATH}/src/github.com/kubeedge/kubeedge/edgesite/; sudo nohup ./edgesite > edgesite.log 2>&1 &"

	CheckCloudcore = "sudo pgrep cloudcore"
	CheckEdgecore  = "sudo pgrep edgecore"
	CheckEdgesite  = "sudo pgrep edgesite"

	CatCloudcoreLog = "cd ${GOPATH}/src/github.com/kubeedge/kubeedge/cloud/; cat cloudcore.log"
	CatEdgecoreLog  = "cd ${GOPATH}/src/github.com/kubeedge/kubeedge/edge/; cat edgecore.log"
	CatEdgeSiteLog  = "cd ${GOPATH}/src/github.com/kubeedge/kubeedge/edgesite/; cat  edgesite.log"

	AppHandler        = "/api/v1/namespaces/default/pods"
	NodeHandler       = "/api/v1/nodes"
	DeploymentHandler = "/apis/apps/v1/namespaces/default/deployments"
)
