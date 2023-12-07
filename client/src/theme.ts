// 1. import `extendTheme` function
import {
  ThemeComponents,
  type ThemeConfig,
  defineStyle,
  extendTheme,
} from "@chakra-ui/react";

// 2. Add your color mode config
const config: ThemeConfig = {
  initialColorMode: "system",
  useSystemColorMode: true,
};

const components: ThemeComponents = {
  Button: {
    sizes: {
      xxs: defineStyle({
        fontSize: "0.65em",
        w: "5",
        h: "5",
      }),
    },
  },
};

// 3. extend the theme
const theme = extendTheme({ config, components });

export default theme;
