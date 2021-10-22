import { generateRelayerRegistry } from "./generateRelayerRegistry.mjs";

export async function createRelayerRegistry({
  chainsProps,
  registryFrom = `/tmp/localnet/registry`,
}) {
  await $`mkdir -p ${registryFrom}`;
  const registry = generateRelayerRegistry(chainsProps);
  await fs.writeFile(`${registryFrom}/registry.yaml`, registry);
}