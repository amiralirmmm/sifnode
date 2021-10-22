import pEvent from "p-event";
import { initChain } from "./initChain.mjs";
import { getChains } from "../utils/getChains.mjs";
import { getChainsProps } from "../utils/getChainsProps.mjs";
import { startChain } from "./startChain.mjs";

export async function buildBinaryNet({
  chainProps: candidateOtherChainProps,
  network,
  rpcInitialPort = 11000,
  p2pInitialPort = 12000,
  pprofInitialPort = 13000,
  home = `/tmp/localnet`,
  registryFrom = `/tmp/localnet/registry`,
}) {
  const chains = getChains({
    rpcInitialPort,
    p2pInitialPort,
    pprofInitialPort,
    home,
  });
  const chainsProps = getChainsProps({ chains, network });
  const sifChainProps = chainsProps.sifchain;
  const otherChainProps = chainsProps[candidateOtherChainProps.chain];

  await $`rm -rf ${home}`;

  await initChain(sifChainProps);
  //   await initChain(otherChainProps);

  const { proc } = await startChain(sifChainProps);
  //   await startChain(otherChainProps);

  // 2b) wait for IBC chains first block to be written
  const asyncIterator = pEvent.iterator(proc.stdout, "data", {
    resolutionEvents: ["finish"],
  });
  for await (let chunk of asyncIterator) {
    if (chunk.includes("indexed block")) break;
  }
}
