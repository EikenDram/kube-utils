// runtime configuration

import type { App, Plugin } from 'vue';

// runtime configuration
export interface RuntimeConfiguration {
    
}

// runtime configuration options
export interface RuntimeConfigurationOptions {
    variables: RuntimeConfiguration
}

export const runtimeConfiguration: Plugin = {
    install: (app: App, options: RuntimeConfigurationOptions) => {
        //Runtime configuration variables can be accessed from injection : `runtimeConfiguration`.
        app.config.globalProperties.$runtimeConfiguration = options.variables

        // They can also be accessed from global property `$runtimeConfiguration`
        app.provide("runtimeConfiguration", options.variables)
    }
}

/**
 * Loads runtime configuration from static file (in /public folder).
 */
export const loadRuntimeConfiguration = async (): Promise<RuntimeConfigurationOptions> => {
    const resp = await fetch('/config.json')
    const value = await resp.json()

    // not type safe but
    return {
        variables: value as RuntimeConfiguration
    } as RuntimeConfigurationOptions
}