package registry

import (
	"fmt"
	"testing"

	mock_clients "github.com/forta-network/forta-node/clients/mocks"
	"github.com/forta-network/forta-node/config"
	mock_store "github.com/forta-network/forta-node/store/mocks"

	"github.com/forta-network/forta-node/services/registry/regtypes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	testScannerAddressStr = "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9"
	testAgentIDStr        = "0x2000000000000000000000000000000000000000000000000000000000000000"
	testAgentRef          = "QmWacxPov5FVCyvnpXroDJ76urakzN4ckpFhhRzpsAkRek"
	testImageRef          = "bafybeide7cspdmxqjcpa3qvrayvfpiix2it4v6mjejjc22q72zbq7rm4re@sha256:cdd4ddccf5e9c740eb4144bcc68e3ea3a056789ec7453e94a6416dcfc80937a4"
	testContainerRegistry = "some.reg.io"
)

var (
	testScannerAddress = common.HexToAddress(testScannerAddressStr)
	testAgentFile      = &regtypes.AgentFile{}
)

// TestSuite runs the test suite.
func TestSuite(t *testing.T) {
	testAgentFile.Manifest.ImageReference = testImageRef

	suite.Run(t, &Suite{})
}

// Suite is a test suite to test the tx node runner implementation.
type Suite struct {
	r *require.Assertions

	registryStore *mock_store.MockRegistryStore
	msgClient     *mock_clients.MockMessageClient

	service *registryService

	suite.Suite
}

// SetupTest sets up the test.
func (s *Suite) SetupTest() {
	s.r = require.New(s.T())
	s.registryStore = mock_store.NewMockRegistryStore(gomock.NewController(s.T()))
	s.msgClient = mock_clients.NewMockMessageClient(gomock.NewController(s.T()))
	s.service = &registryService{
		scannerAddress: testScannerAddress,
		registryStore:  s.registryStore,
	}
	s.service.cfg.Registry.ContainerRegistry = testContainerRegistry
}

type agentConfigs []*config.AgentConfig

func (ac agentConfigs) Matches(x interface{}) bool {
	acx, ok := x.([]*config.AgentConfig)
	if !ok {
		return false
	}

	if len(ac) != len(acx) {
		return false
	}

	for i, agent1 := range ac {
		agent2 := acx[i]
		if !(agent1.ID == agent2.ID && agent1.Image == agent2.Image) {
			return false
		}
	}
	return true
}

func (ac agentConfigs) String() string {
	return fmt.Sprintf("%+v", ([]*config.AgentConfig)(ac))
}

func (s *Suite) TestReturnList() {
	configs := (agentConfigs)([]*config.AgentConfig{
		{
			ID:    testAgentIDStr,
			Image: fmt.Sprintf("%s/%s", testContainerRegistry, testImageRef),
		},
	})

	s.registryStore.EXPECT().GetAgentsIfChanged(s.service.scannerAddress.Hex()).Return(configs, true, nil)

	received, err := s.service.GetLatestBots()
	s.r.NoError(err)
	s.r.Len(received, 1)
	s.r.Equal(configs[0].ID, received[0].ID)
}

func (s *Suite) TestReturnListEvenIfNoChanges() {
	configs := (agentConfigs)([]*config.AgentConfig{
		{
			ID:    testAgentIDStr,
			Image: fmt.Sprintf("%s/%s", testContainerRegistry, testImageRef),
		},
	})

	// first refresh
	s.registryStore.EXPECT().GetAgentsIfChanged(s.service.scannerAddress.Hex()).Return(configs, true, nil)
	received, err := s.service.GetLatestBots()
	s.r.NoError(err)
	s.r.Len(received, 1)
	s.r.Equal(configs[0].ID, received[0].ID)

	// second refresh
	s.registryStore.EXPECT().GetAgentsIfChanged(s.service.scannerAddress.Hex()).Return(nil, false, nil)
	received, err = s.service.GetLatestBots()
	s.r.NoError(err)
	s.r.Len(received, 1)
	s.r.Equal(configs[0].ID, received[0].ID)
}
