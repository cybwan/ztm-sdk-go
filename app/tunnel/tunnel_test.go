package tunnel

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"

	ztm "github.com/cybwan/ztm-sdk-go"
	"github.com/cybwan/ztm-sdk-go/local"
)

func TestListOutbounds(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockLocalSpec := local.NewMockZtmLocalSpec(mockCtrl)
	mockLocalSpec.EXPECT().GetLocalUser().AnyTimes().Return("root")
	mockLocalSpec.EXPECT().GetLocalMesh().AnyTimes().Return("k8s-mesh")

	agent := os.Getenv("CTR_AGENT")
	if strings.EqualFold(agent, "office") {
		mockLocalSpec.EXPECT().GetLocalAgent().AnyTimes().Return("127.0.0.1:7772")
	} else {
		mockLocalSpec.EXPECT().GetLocalAgent().AnyTimes().Return("127.0.0.1:7771")
	}

	hostIP := os.Getenv("MY_HOST_IP")
	if len(hostIP) == 0 {
		hostIP = "127.0.0.1"
	}
	mockLocalSpec.EXPECT().GetLocalHostAddr().AnyTimes().Return(hostIP)

	agentClient := ztm.NewAgentClient(mockLocalSpec.GetLocalAgent(), false)

	if localEndpoint, err := agentClient.LocalEndpoint(mockLocalSpec.GetLocalMesh()); err == nil {
		mockLocalSpec.EXPECT().GetLocalEndpointId().AnyTimes().Return(localEndpoint.UUID)
	}

	tunnelClient := Client{
		OutboundClient: agentClient.OutboundClient,
		InboundClient:  agentClient.InboundClient,
	}

	fmt.Println(mockLocalSpec.GetLocalEndpointId())

	outbounds, _ := tunnelClient.ListOutbounds(mockLocalSpec.GetLocalMesh(), mockLocalSpec.GetLocalEndpointId(), ztm.ZTM)
	bytes, _ := json.MarshalIndent(outbounds, "", " ")
	fmt.Println(string(bytes))
}
