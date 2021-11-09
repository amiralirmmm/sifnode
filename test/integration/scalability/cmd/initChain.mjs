#!/usr/bin/env zx

import { initChain } from "../lib/initChain.mjs";
import { arg } from "../utils/arg.mjs";
import { getChainProps } from "../utils/getChainProps.mjs";

const args = arg(
  {
    "--chain": String,
    "--network": String,
    "--node": String,
    "--chain-id": String,
    "--binary": String,
    "--name": String,
    "--amount": Number,
    "--denom": String,
    "--home": String,
  },
  `
Usage:

  yarn initChain [options]

Initiate a new chain locally based on an existing remote chain.

Options:

--chain     Select a predifined chain in chains.json
--network   Select a predifined network in chains.json
--node      Node address
--chain-id  Chain ID
--binary    Binary name of the chain
--name      Account name
--amount    Amount to send to receiver account
--denom     Chain denom
--home      Directory for config and data
`
);

const chain = args["--chain"] || undefined;
const network = args["--network"] || undefined;
const node = args["--node"] || undefined;
const chainId = args["--chain-id"] || undefined;
const binary = args["--binary"] || undefined;
const name = args["--name"] || undefined;
const amount = args["--amount"] || undefined;
const denom = args["--denom"] || undefined;
const home = args["--home"] || undefined;

const chainProps = getChainProps({
  chain,
  network,
  node,
  chainId,
  binary,
  name,
  amount,
  denom,
  home,
});
await initChain({
  ...chainProps,
});