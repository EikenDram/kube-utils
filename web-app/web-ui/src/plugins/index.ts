/**
 * plugins/index.ts
 *
 * Automatically included in `./src/main.ts`
 */

// Plugins
import { loadFonts } from "./webfontloader";
import vuetify from "./vuetify";
import i18n from "./i18n";
import router from "../router";
import axios from "./axios";

// Types
import type { App } from "vue";

export function registerPlugins(app: App) {
  loadFonts();

  app.use(i18n);
  app.use(axios).use(vuetify).use(router);
}
