<template>
  <v-container fluid>

    <v-stepper v-model="step" :items="items" show-actions>
      <template v-slot:item.1>
        <h3 class="text-h6">Order</h3>

        <br>

        <v-file-input accept="*" :label="$t('keycloak.file')"></v-file-input>
      </template>

      <template v-slot:item.2>
        <h3 class="text-h6">{{ $t('keycloak.operation') }}</h3>

        <br>

        <v-radio-group v-model="operation">
          <v-radio :label="$t('keycloak.update')" :value="Operation.update"></v-radio>
          <v-radio :label="$t('keycloak.create')" :value="Operation.create"></v-radio>
          <v-radio :label="$t('keycloak.delete')" :value="Operation.delete"></v-radio>
        </v-radio-group>
      </template>

      <template v-slot:item.3>
        <h3 class="text-h6">Confirm</h3>

        <br>

        <v-sheet border>
          Content
        </v-sheet>
      </template>
    </v-stepper>
  </v-container>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

enum Operation {
  update,
  create,
  delete,
}

const step = ref(1)

const operation = ref<Operation>(Operation.update)

const items = [
  'keycloak.uploadFile',
  'keycloak.selectOperation',
  'keycloak.selectFormat',
  'keycloak.confirm',
  'keycloak.result',
]
</script>