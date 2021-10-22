#!/usr/bin/env zx

import { startAllChains } from "../lib/startAllChains.mjs";
import { arg } from "../utils/arg.mjs";
import { getChainProps } from "../utils/getChainProps.mjs";

const args = arg(
  {
    "--network": String,
    "--rpcInitialPort": Number,
    "--p2pInitialPort": Number,
    "--pprofInitialPort": Number,
    "--home": String,
  },
  `
Usage:

  yarn startAllChains [options]

Start all the chains locally.

Options:

--network               Select a predifined network in chains.json
--rpcInitialPort        Initial RPC port number
--p2pInitialPort        Initial P2P port number
--pprofInitialPort      Initial pprof port number
--home                  Global directory for config and data of initiated chains
`
);

const network = args["--network"] || undefined;
const rpcInitialPort = args["--rpcInitialPort"] || undefined;
const p2pInitialPort = args["--p2pInitialPort"] || undefined;
const pprofInitialPort = args["--pprofInitialPort"] || undefined;
const home = args["--home"] || undefined;

const chainProps = getChainProps({
  network,
  rpcInitialPort,
  p2pInitialPort,
  pprofInitialPort,
  home,
});
await startAllChains({
  ...chainProps,
});