package cli

import (
	
    "github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/b9lab/toll-road/x/tollroad/types"
)

func CmdCreateUserVault() *cobra.Command {
    cmd := &cobra.Command{
		Use:   "create-user-vault [index] [road-operator-index] [token] [balance]",
		Short: "Create a new UserVault",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            // Get indexes
         indexIndex := args[0]
        
            // Get value arguments
		 argRoadOperatorIndex := args[1]
		 argToken := args[2]
		 argBalance := args[3]
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUserVault(
			    clientCtx.GetFromAddress().String(),
			    indexIndex,
                argRoadOperatorIndex,
			    argToken,
			    argBalance,
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

func CmdUpdateUserVault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-user-vault [index] [road-operator-index] [token] [balance]",
		Short: "Update a UserVault",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            // Get indexes
         indexIndex := args[0]
        
            // Get value arguments
		 argRoadOperatorIndex := args[1]
		 argToken := args[2]
		 argBalance := args[3]
		
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateUserVault(
			    clientCtx.GetFromAddress().String(),
			    indexIndex,
                argRoadOperatorIndex,
                argToken,
                argBalance,
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

func CmdDeleteUserVault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-user-vault [index]",
		Short: "Delete a UserVault",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
             indexIndex := args[0]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteUserVault(
			    clientCtx.GetFromAddress().String(),
			    indexIndex,
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