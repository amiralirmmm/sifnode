package cli_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"testing"

	testutils "github.com/Sifchain/sifnode/x/dispensation/client/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

const (
	AccountAddressPrefix = "sif"
)

var (
	AccountPubKeyPrefix    = AccountAddressPrefix + "pub"
	ValidatorAddressPrefix = AccountAddressPrefix + "valoper"
	ValidatorPubKeyPrefix  = AccountAddressPrefix + "valoperpub"
	ConsNodeAddressPrefix  = AccountAddressPrefix + "valcons"
	ConsNodePubKeyPrefix   = AccountAddressPrefix + "valconspub"
)

func SetConfig(seal bool) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	if seal {
		config.Seal()
	}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()
	//genesisState := cfg.GenesisState
	//cfg.NumValidators = 1
	SetConfig(true)
	//var bankGenesis types.GenesisState
	//s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &bankGenesis))
	//
	//bankGenesis.DenomMetadata = []types.Metadata{
	//	{
	//		Description: "The native staking token ",
	//		DenomUnits: []*types.DenomUnit{
	//			{
	//				Denom:    "rowan",
	//				Exponent: 18,
	//				Aliases:  []string{"RWN"},
	//			},
	//		},
	//		Base:    "rowan",
	//		Display: "rowan",
	//	},
	//}
	//
	//bankGenesisBz, err := cfg.Codec.MarshalJSON(&bankGenesis)
	//s.Require().NoError(err)
	//genesisState[types.ModuleName] = bankGenesisBz
	//cfg.GenesisState = genesisState

	s.cfg = cfg
	s.cfg.InterfaceRegistry.RegisterInterface("sifnode.dispensation.v1.MsgCreateDistribution", (*sdk.Msg)(nil))
	s.network = network.New(s.T(), cfg)
	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestCreateDispensationTxCmd() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx
	args := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
	}
	bz, err := testutils.MsgCreateDispensationExec(clientCtx, "ValidatorSubsidy", "output.json", val.Address.String(), args...)
	s.Require().NoError(err)
	respType := &sdk.TxResponse{}
	s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	fmt.Println(respType)
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
