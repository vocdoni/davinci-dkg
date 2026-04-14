import { extendTheme, type ThemeConfig } from '@chakra-ui/react';

const config: ThemeConfig = {
  initialColorMode: 'dark',
  useSystemColorMode: true,
};

export const theme = extendTheme({
  config,
  fonts: {
    heading: '"Inter", -apple-system, system-ui, sans-serif',
    body: '"Inter", -apple-system, system-ui, sans-serif',
    mono: '"JetBrains Mono", "Menlo", monospace',
  },
  styles: {
    global: {
      'html, body, #root': {
        minHeight: '100%',
      },
    },
  },
});
