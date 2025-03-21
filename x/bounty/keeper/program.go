package keeper

import (
	"crypto/sha256"
	"encoding/hex"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/shentufoundation/shentu/v2/x/bounty/types"
)

func (k Keeper) GetProgram(ctx sdk.Context, id string) (types.Program, bool) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	pBz := store.Get(types.GetProgramKey(id))
	if pBz == nil {
		return types.Program{}, false
	}

	var program types.Program
	k.cdc.MustUnmarshal(pBz, &program)
	return program, true
}

func (k Keeper) GetAllPrograms(ctx sdk.Context) []types.Program {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	var programs []types.Program
	var program types.Program

	iterator := storetypes.KVStorePrefixIterator(store, types.ProgramKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		k.cdc.MustUnmarshal(iterator.Value(), &program)
		programs = append(programs, program)
	}
	return programs
}

func (k Keeper) SetProgram(ctx sdk.Context, program types.Program) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := k.cdc.MustMarshal(&program)
	store.Set(types.GetProgramKey(program.ProgramId), bz)
}

func (k Keeper) ActivateProgram(ctx sdk.Context, pid string, caller sdk.AccAddress) error {
	program, found := k.GetProgram(ctx, pid)
	if !found {
		return types.ErrProgramNotExists
	}

	// Check if the program is already active
	if program.Status == types.ProgramStatusActive {
		return types.ErrProgramAlreadyActive
	}

	// Check the permissions. Only the bounty cert address can operate.
	if !k.certKeeper.IsBountyAdmin(ctx, caller) {
		return types.ErrProgramOperatorNotAllowed
	}

	program.Status = types.ProgramStatusActive
	k.SetProgram(ctx, program)
	return nil
}

func (k Keeper) CloseProgram(ctx sdk.Context, pid string, caller sdk.AccAddress) error {
	program, found := k.GetProgram(ctx, pid)
	if !found {
		return types.ErrProgramNotExists
	}

	// Check if the program is already closed
	if program.Status == types.ProgramStatusClosed {
		return types.ErrProgramAlreadyClosed
	}

	// The program cannot be closed
	// There are 3 finding states: FindingStatusSubmitted FindingStatusActive FindingStatusConfirmed
	fidsList, err := k.GetPidFindingIDList(ctx, pid)
	if err != nil {
		return err
	}
	for _, fid := range fidsList {
		finding, found := k.GetFinding(ctx, fid)
		if !found {
			return types.ErrFindingNotExists
		}
		if finding.Status == types.FindingStatusSubmitted ||
			finding.Status == types.FindingStatusActive ||
			finding.Status == types.FindingStatusConfirmed {
			return types.ErrProgramCloseNotAllowed
		}
	}

	// Check the permissions. Only the admin of the program or bounty cert address can operate.
	if program.AdminAddress != caller.String() && !k.certKeeper.IsBountyAdmin(ctx, caller) {
		return types.ErrProgramOperatorNotAllowed
	}

	// Close the program and update its status in the store
	program.Status = types.ProgramStatusClosed
	k.SetProgram(ctx, program)
	return nil
}

func (k Keeper) GetProgramFingerprintHash(program *types.Program) string {
	programFingerprint := &types.ProgramFingerprint{
		ProgramId:    program.ProgramId,
		Name:         program.Name,
		Detail:       program.Detail,
		AdminAddress: program.AdminAddress,
		Status:       program.Status,
	}

	bz := k.cdc.MustMarshal(programFingerprint)
	hash := sha256.Sum256(bz)
	return hex.EncodeToString(hash[:])
}
