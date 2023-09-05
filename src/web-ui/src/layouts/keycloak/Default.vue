<template>
  <v-app>
    <keycloak-bar />

    <keycloak-navigation />

    <keycloak-view />
  </v-app>
</template>

<script lang="ts" setup>
import KeycloakBar from "./AppBar.vue";
import KeycloakNavigation from "./Navigation.vue";
import KeycloakView from "./View.vue";
</script>
