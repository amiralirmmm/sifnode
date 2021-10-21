package testutil

import (
	dispensationcli "github.com/Sifchain/sifnode/x/dispensation/client/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

func MsgCreateDispensationExec(clientCtx client.Context, distributionType, outputsPath, authorizedRunner string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{distributionType, outputsPath, authorizedRunner}
	args = append(args, extraArgs...)

	return clitestutil.ExecTestCLICmd(clientCtx, dispensationcli.GetCmdCreate(), args)
}

//
//// serviceMsgClientConn is an instance of grpc.ClientConn that is used to test building
//// transactions with MsgClient's. It is intended to be replaced by the work in
//// https://github.com/cosmos/cosmos-sdk/issues/7541 when that is ready.
//type serviceMsgClientConn struct {
//	msgs []sdk.Msg
//}
//
//func (t *serviceMsgClientConn) Invoke(_ context.Context, method string, args, _ interface{}, _ ...grpc.CallOption) error {
//	req, ok := args.(sdk.MsgRequest)
//	if !ok {
//		return fmt.Errorf("%T should implement %T", args, (*sdk.MsgRequest)(nil))
//	}
//
//	err := req.ValidateBasic()
//	if err != nil {
//		return err
//	}
//
//	t.msgs = append(t.msgs, sdk.ServiceMsg{
//		MethodName: method,
//		Request:    req,
//	})
//
//	return nil
//}
//
//func (t *serviceMsgClientConn) NewStream(context.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
//	return nil, fmt.Errorf("not supported")
//}
//
//var _ gogogrpc.ClientConn = &serviceMsgClientConn{}
//
//// newSendTxMsgServiceCmd is just for the purpose of testing ServiceMsg's in an end-to-end case. It is effectively
//// NewSendTxCmd but using MsgClient.
//func newCreateTxMsgServiceCmd() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "create [DistributionType] [Output JSON File Path] [AuthorizedRunner]",
//		Short: "Create new distribution",
//		RunE: func(cmd *cobra.Command, args []string) error {
//			clientCtx, err := client.GetClientTxContext(cmd)
//			if err != nil {
//				return err
//			}
//			err = cobra.ExactArgs(3)(cmd, args)
//			if err != nil {
//				return err
//			}
//			distributionType, ok := dispensation.GetDistributionTypeFromShortString(args[0])
//			if !ok {
//				return fmt.Errorf("invalid distribution Type %s: Types supported [Airdrop/LiquidityMining/ValidatorSubsidy]", args[2])
//			}
//			outputList, err := dispensationUtils.ParseOutput(args[1])
//			if err != nil {
//				return err
//			}
//			msg := dispensation.NewMsgCreateDistribution(clientCtx.GetFromAddress(), distributionType, outputList, args[2])
//			svcMsgClientConn := &serviceMsgClientConn{}
//			bankMsgClient := dispensation.NewMsgClient(svcMsgClientConn)
//			_, err = bankMsgClient.CreateDistribution(context.Background(), &msg)
//			if err != nil {
//				return err
//			}
//
//			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.msgs...)
//		},
//	}
//
//	flags.AddTxFlagsToCmd(cmd)
//
//	return cmd
//}
//
//// ServiceMsgSendExec is a temporary method to test Msg services in CLI using
//// x/bank's Msg/Send service. After https://github.com/cosmos/cosmos-sdk/issues/7541
//// is merged, this method should be removed, and we should prefer MsgSendExec
//// instead.
//func ServiceMsgCreateDispensationExec(clientCtx client.Context, from, to, amount fmt.Stringer, extraArgs ...string) (testutil.BufferWriter, error) {
//	args := []string{from.String(), to.String(), amount.String()}
//	args = append(args, extraArgs...)
//
//	return clitestutil.ExecTestCLICmd(clientCtx, newCreateTxMsgServiceCmd(), args)
//}
