<template>
  <v-navigation-drawer permanent>
    <v-list>
      <v-list-item
        prepend-icon="fa-solid fa-list"
        :title="$t('keycloak.list')"
        to="/keycloak/list"
      ></v-list-item>
      <v-list-item
        prepend-icon="fa-solid fa-file-import"
        :title="$t('keycloak.processFile')"
        to="/keycloak/process"
      ></v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script lang="ts" setup>
//
</script>
