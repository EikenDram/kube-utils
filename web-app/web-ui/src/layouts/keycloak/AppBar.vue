<template>
  <v-app-bar flat>
    <v-app-bar-title>
      <router-link to="/">
        <v-icon color="black" class="ml-2 mr-2" icon="fa-solid fa-cubes" />
      </router-link>

      {{ $t("keycloak.name") }}
    </v-app-bar-title>
  </v-app-bar>
</template>

<script lang="ts" setup>
//
</script>
