# **accounts add hardware-wallet**

`accounts add hardware-wallet [-ledger-live] <wallet_type> <name> <address>`

**Alias:** `hw`

Adds a hardware wallet account. Uses the Ledger Live derivation path if **-ledger-live** is specified.

- **wallet_type**: the hardware wallet type; can be `ledger`, `trezor`, or `smartcard`.
- **name**: the name of the account to add.
- **address**: the wallet address to add. Must match one derived from the wallet.
