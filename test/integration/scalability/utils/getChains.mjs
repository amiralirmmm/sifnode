const defaultChains = require("../config/chains.json");

export function getChains({
  chains = defaultChains,
  rpcInitialPort = 11000,
  p2pInitialPort = 12000,
  pprofInitialPort = 13000,
  home = `/tmp/localnet`,
}) {
  const newChains = { ...chains };

  Object.keys(newChains).forEach((chain, index) => {
    newChains[chain] = {
      ...newChains[chain],
      rpcPort: rpcInitialPort + index,
      p2pPort: p2pInitialPort + index,
      pprofPort: pprofInitialPort + index,
      home: `${home}/${chain}/${newChains[chain].chainId}`,
    };
  });

  return newChains;
}