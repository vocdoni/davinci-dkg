import { getAddress, isAddress } from 'viem'

/**
 * EIP-55 checksum, or the input unchanged when it isn't an address.
 *
 * Every address the explorer displays goes through this: an event log gives
 * lowercase addresses, and showing them next to checksummed ones from a
 * contract read makes the same operator look like two.
 */
export function checksum(value: string): string {
  return isAddress(value, { strict: false }) ? getAddress(value) : value
}
