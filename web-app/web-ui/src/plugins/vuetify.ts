/**
 * plugins/vuetify.ts
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
//import '@mdi/font/css/materialdesignicons.css'
import "@fortawesome/fontawesome-free/css/all.css"; // Ensure your project is capable of handling css files
import { aliases, fa } from "vuetify/iconsets/fa";
import "vuetify/styles";

// Composables
// vuetify

import { createVuetify } from "vuetify";

import * as components from "vuetify/components";
import * as labsComponents from "vuetify/labs/components";

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  icons: {
    defaultSet: "fa",
    aliases,
    sets: {
      fa,
    },
  },
  theme: {
    themes: {
      light: {
        colors: {
          primary: "#1867C0",
          secondary: "#5CBBF6",
        },
      },
    },
  },
  components: {
    ...components,
    ...labsComponents,
  },
});
