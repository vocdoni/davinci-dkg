// Package txmanager provides a minimal nonce-aware transaction manager on
// top of go-ethereum's ethclient.
//
// It signs EIP-1559 transactions with a single key, tracks pending nonces,
// retries stuck transactions with increased tip/fee, and waits for receipts.
// It is intentionally narrow in scope: a single sender per instance, no
// mempool awareness, and no reorg handling beyond what TransactionReceipt
// exposes. It is sufficient for DKG nodes and the testnet runner, not for
// general-purpose wallet software.
package txmanager
