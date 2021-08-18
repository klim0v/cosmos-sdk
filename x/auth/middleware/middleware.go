package middleware

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
)

// ComposeTxMiddleware compose multiple middlewares on top of a TxHandler. Last
// middleware is the outermost middleware.
func ComposeTxMiddleware(txHandler tx.TxHandler, middlewares ...tx.TxMiddleware) tx.TxHandler {
	for _, m := range middlewares {
		txHandler = m(txHandler)
	}

	return txHandler
}

type TxHandlerOptions struct {
	Debug bool
	// IndexEvents defines the set of events in the form {eventType}.{attributeKey},
	// which informs Tendermint what to index. If empty, all events will be indexed.
	IndexEvents map[string]struct{}

	LegacyRouter     sdk.Router
	MsgServiceRouter *MsgServiceRouter

	LegacyAnteHandler sdk.AnteHandler
}

// NewDefaultTxHandler defines a TxHandler middleware stacks that should work
// for most applications.
func NewDefaultTxHandler(options TxHandlerOptions) tx.TxHandler {
	return ComposeTxMiddleware(
		NewRunMsgsTxHandler(options.MsgServiceRouter, options.LegacyRouter),
		newLegacyAnteMiddleware(options.LegacyAnteHandler),
		// Choose which events to index in Tendermint. Make sure no events are
		// emitted outside of this middleware.
		NewIndexEventsTxMiddleware(options.IndexEvents),
		// Recover from panics. Panics outside of this middleware won't be
		// caught, be careful!
		NewRecoveryTxMiddleware(),
		// Convert errors into ABCI responses.
		NewErrorTxMiddleware(options.Debug),
		// Set a new GasMeter on sdk.Context.
		//
		// Make sure the Gas middleware is outside of all other middlewares
		// that reads the GasMeter. In our case, the Error middleware reads
		// the GasMeter to populate GasInfo.
		NewGasTxMiddleware(),
	)
}