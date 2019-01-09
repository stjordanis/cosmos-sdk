package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/stretchr/testify/require"
)

func TestWithdrawValidatorCommission(t *testing.T) {
	ctx, ak, keeper, _, _ := CreateTestInputDefault(t, false, 1000)

	// set zero outstanding rewards
	keeper.SetOutstandingRewards(ctx, types.OutstandingRewards{})

	// check initial balance
	balance := ak.GetAccount(ctx, sdk.AccAddress(valOpAddr3)).GetCoins()
	require.Equal(t, balance, sdk.Coins{
		{"stake", sdk.NewInt(1000)},
	})

	// set commission
	keeper.SetValidatorAccumulatedCommission(ctx, valOpAddr3, sdk.DecCoins{
		{"mytoken", sdk.NewDec(5).Quo(sdk.NewDec(4))},
		{"stake", sdk.NewDec(3).Quo(sdk.NewDec(2))},
	})

	// withdraw commission
	keeper.WithdrawValidatorCommission(ctx, valOpAddr3)

	// check balance increase
	balance = ak.GetAccount(ctx, sdk.AccAddress(valOpAddr3)).GetCoins()
	require.Equal(t, balance, sdk.Coins{
		{"mytoken", sdk.NewInt(1)},
		{"stake", sdk.NewInt(1001)},
	})

	// check remainder
	remainder := keeper.GetValidatorAccumulatedCommission(ctx, valOpAddr3)
	require.Equal(t, remainder, sdk.DecCoins{
		{"mytoken", sdk.NewDec(1).Quo(sdk.NewDec(4))},
		{"stake", sdk.NewDec(1).Quo(sdk.NewDec(2))},
	})

	require.True(t, true)
}
