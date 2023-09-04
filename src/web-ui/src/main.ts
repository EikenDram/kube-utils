/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'

// Configuration
import {loadRuntimeConfiguration, runtimeConfiguration} from "@/plugins/runtimeConfiguration";

loadRuntimeConfiguration()
    .then((runtimeConfigurationOptions) => {
        const app = createApp(App)

        app.use(runtimeConfiguration, runtimeConfigurationOptions)

        registerPlugins(app)

        app.mount('#app')
    });

