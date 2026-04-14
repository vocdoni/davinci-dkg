import 'vitest';

// Augment Vitest's ProvidedContext with the keys provided by globalSetup.
declare module 'vitest' {
  export interface ProvidedContext {
    integrationEnabled: boolean;
    rpcUrl:             string;
    managerAddress:     string;
    registryAddress:    string;
    addressesFile:      string;
  }
}
