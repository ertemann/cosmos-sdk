package cli

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"cosmossdk.io/x/upgrade/plan"
	"cosmossdk.io/x/upgrade/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	// Deprecated: only used for v1beta1 legacy proposals.
	FlagUpgradeHeight = "upgrade-height"
	// Deprecated: only used for v1beta1 legacy proposals.
	FlagUpgradeInfo = "upgrade-info"
	FlagNoValidate  = "no-validate"
	FlagDaemonName  = "daemon-name"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Upgrade transaction subcommands",
	}

	return cmd
}

// NewCmdSubmitUpgradeProposal implements a command handler for submitting a software upgrade proposal transaction.
func NewCmdSubmitUpgradeProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "software-upgrade [name] (--upgrade-height [height]) (--upgrade-info [info]) [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a software upgrade proposal",
		Long: "Submit a software upgrade along with an initial deposit.\n" +
			"Please specify a unique name and height for the upgrade to take effect.\n" +
			"You may include info to reference a binary download link, in a format compatible with: https://docs.cosmos.network/main/tooling/cosmovisor",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposal, err := cli.ReadGovPropFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			name := args[0]
			p, err := parsePlan(cmd.Flags(), name)
			if err != nil {
				return err
			}

			noValidate, err := cmd.Flags().GetBool(FlagNoValidate)
			if err != nil {
				return err
			}

			if !noValidate {
				var daemonName string
				if daemonName, err = cmd.Flags().GetString(FlagDaemonName); err != nil {
					return err
				}

				var planInfo *plan.Info
				if planInfo, err = plan.ParseInfo(p.Info); err != nil {
					return err
				}
				if err = planInfo.ValidateFull(daemonName); err != nil {
					return err
				}
			}

			proposal.SetMsgs([]sdk.Msg{
				&types.MsgSoftwareUpgrade{
					Plan: p,
				},
			})

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), proposal)
		},
	}

	cmd.Flags().Int64(FlagUpgradeHeight, 0, "The height at which the upgrade must happen")
	cmd.Flags().String(FlagUpgradeInfo, "", "Info for the upgrade plan such as new version download urls, etc.")
	cmd.Flags().Bool(FlagNoValidate, false, "Skip validation of the upgrade info")
	cmd.Flags().String(FlagDaemonName, getDefaultDaemonName(), "The name of the executable being upgraded (for upgrade-info validation). Default is the DAEMON_NAME env var if set, or else this executable")

	// add common proposal flags
	cli.AddGovPropFlagsToCmd(cmd)

	return cmd
}

// NewCmdSubmitLegacyCancelUpgradeProposal implements a command handler for submitting a software upgrade cancel proposal transaction.
// Deprecated: please use NewCmdSubmitCancelUpgradeProposal instead.
func NewCmdSubmitLegacyCancelUpgradeProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-software-upgrade [flags]",
		Args:  cobra.ExactArgs(0),
		Short: "Cancel the current software upgrade proposal",
		Long:  "Cancel a software upgrade along with an initial deposit.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription) //nolint:staticcheck // we are intentionally using a deprecated flag here.
			if err != nil {
				return err
			}

			content := types.NewCancelSoftwareUpgradeProposal(title, description)

			msg, err := v1beta1.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

// getDefaultDaemonName gets the default name to use for the daemon.
// If a DAEMON_NAME env var is set, that is used.
// Otherwise, the last part of the currently running executable is used.
func getDefaultDaemonName() string {
	// DAEMON_NAME is specifically used here to correspond with the Cosmovisor setup env vars.
	name := os.Getenv("DAEMON_NAME")
	if len(name) == 0 {
		_, name = filepath.Split(os.Args[0])
	}
	return name
}
