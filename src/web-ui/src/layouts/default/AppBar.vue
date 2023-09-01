<template>
  <v-app-bar flat>
    <v-app-bar-title>
      <v-icon icon="mdi-circle-slice-4" />

      <a to="/">{{ locale.AppName }}</a>
      
    </v-app-bar-title>
  </v-app-bar>
</template>

<script lang="ts" setup>
import {inject} from 'vue'
import type {RuntimeConfiguration} from "@/plugins/runtimeConfiguration";

const rc = inject('runtimeConfiguration') as RuntimeConfiguration
const locale = rc.locale
</script>
