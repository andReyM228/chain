package cli

import (
	"strconv"

	"github.com/andReyM228/one/x/core/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [amount] [address] [denom]",
		Short: "Broadcast message withdraw",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argAddress := args[1]
			argDenom := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdraw(
				clientCtx.GetFromAddress().String(),
				argAmount,
				argAddress,
				argDenom,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
