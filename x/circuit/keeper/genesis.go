package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/circuit/types"
)

func (k *Keeper) ExportGenesis(ctx sdk.Context) (data *types.GenesisState) {
	var permissions []*types.GenesisAccountPermissions
	var disabledMsgs []string

	k.IteratePermissions(ctx, func(address []byte, perm types.Permissions) (stop bool) {
		// Convert the Permissions struct to a GenesisAccountPermissions struct
		// and add it to the permissions slice
		permissions = append(permissions, &types.GenesisAccountPermissions{
			Address:     string(address),
			Permissions: &perm,
		})
		return false
	})

	k.IterateDisableLists(ctx, func(address []byte, perm types.Permissions) (stop bool) {
		for _, url := range perm.LimitTypeUrls {
			disabledMsgs = append(disabledMsgs, url)
		}
		return false
	})

	return &types.GenesisState{
		AccountPermissions: permissions,
		DisabledTypeUrls:   disabledMsgs,
	}
}

// InitGenesis initializes the bank module's state from a given genesis state.
func (k *Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {

	for _, accounts := range genState.AccountPermissions {
		// Set the permissions for the account
		k.SetPermissions(ctx, accounts.Address, accounts.Permissions)
	}
	for _, url := range genState.DisabledTypeUrls {
		// Set the disabled type urls
		k.DisableMsg(ctx, url)
	}
}