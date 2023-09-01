/**
 * plugins/index.ts
 *
 * Automatically included in `./src/main.ts`
 */

// Plugins
import { loadFonts } from './webfontloader'
import vuetify from './vuetify'
import router from '../router'
import { loadRuntimeConfiguration, runtimeConfiguration } from "./runtimeConfiguration";
import { createI18n } from 'vue-i18n'

import en from '../locales/en.json'

const i18n = createI18n({
    locale: 'en',
    messages: {
      en,
    }
  })

const runtimeConfigurationOptions = await loadRuntimeConfiguration()

// Types
import type { App } from 'vue'

export function registerPlugins (app: App) {
  loadFonts()
  app
    .use(i18n)
    .use(vuetify)
    .use(router)
    .use(runtimeConfiguration, runtimeConfigurationOptions)
}
