<template>
  <v-container class="fill-height">
    <v-responsive class="align-center text-center fill-height">

      <h1 class="text-h2 font-weight-bold">{{ $t('app.name') }}</h1>

      <div class="text-body-2 font-weight-light mb-n1">{{ $t('app.desc') }}</div>

      <div class="py-14" />

      <v-row class="d-flex align-center justify-center">


        <v-col cols="auto">
          <v-btn color="primary" to="/keycloak" min-width="228" size="x-large" variant="flat">
            <v-icon icon="fa-solid fa-user-gear" size="large" start />

            Keycloak
          </v-btn>
        </v-col>


      </v-row>
    </v-responsive>
  </v-container>
</template>

<script lang="ts" setup>
import { inject } from 'vue'
import type { RuntimeConfiguration } from "@/plugins/runtimeConfiguration";

const rc = inject('runtimeConfiguration') as RuntimeConfiguration

</script>
