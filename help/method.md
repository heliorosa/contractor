# **method (m)**

`method [-nowait] [-network <network>] <account> <contract> <method_name> [method_args ...]`

**Alias:** `m`

Calls **method_name** on a **contract** using the specified **account**, passing any additional **method_args**.

If the method is `view` or `pure`, the result is displayed. Otherwise, a transaction is sent and awaited — unless **-nowait** is used.

If the **contract** is provided in the format `address:abi_file`, the **-network** flag is required and must be either `@network_name` (preconfigured) or a provider URL.

- **account** must be preconfigured. It can be in the form `@account_name` or a raw address.
- **contract** can be `@contract_name` (preconfigured) or in the format `address:abi_file`.
- **method_name** must exist in the contract’s ABI.
- **-network** (optional): required when using the `address:abi_file` contract format. Accepts `@network_name` or a provider URL.
- **-nowait** (optional): if specified, the transaction is sent but not awaited for mining.
