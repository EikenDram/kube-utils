<template>
    <v-container class="fill-height">
      <v-responsive class="align-center text-center fill-height">
        <v-row class="d-flex align-center justify-center">
             Keycloak management
        </v-row>
      </v-responsive>
    </v-container>
  </template>
  
  <script lang="ts" setup>
  import { inject } from 'vue'
  import type {RuntimeConfiguration} from "@/plugins/runtimeConfiguration";
  
  const rc = inject('runtimeConfiguration') as RuntimeConfiguration
  const locale = rc.locale
  </script>
  