<template>
  <v-container fluid>
    <v-stepper v-model="step" alt-labels hide-actions>
      <template v-slot:default>
        <v-stepper-header>
          <v-stepper-item
            :complete="step > 1"
            :title="$t('keycloak.selectFile')"
            value="1"
          ></v-stepper-item>

          <v-divider></v-divider>

          <v-stepper-item
            :complete="step > 2"
            :title="$t('keycloak.selectOperation')"
            value="2"
          ></v-stepper-item>

          <v-divider></v-divider>

          <v-stepper-item
            :complete="step > 3"
            :title="$t('keycloak.selectFormat')"
            value="3"
          ></v-stepper-item>

          <v-divider></v-divider>

          <v-stepper-item
            :complete="step > 4"
            :title="$t('keycloak.preview')"
            value="4"
          ></v-stepper-item>

          <v-divider></v-divider>

          <v-stepper-item
            :complete="step > 5"
            :title="$t('keycloak.processing')"
            value="5"
          ></v-stepper-item>
        </v-stepper-header>

        <v-stepper-window>
          <v-stepper-window-item value="1">
            <v-card :title="$t('keycloak.selectFile')" flat>
              <v-file-input
                accept="*"
                :label="$t('keycloak.file')"
              ></v-file-input>

              <v-card-actions>
                <v-btn variant="flat" color="primary" @click="uploadFile">
                  {{ $t("keycloak.upload") }}
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-stepper-window-item>

          <v-stepper-window-item value="2">
            <v-card :title="$t('keycloak.selectOperation')" flat>
              <v-radio-group v-model="operation">
                <v-radio
                  :label="$t('keycloak.operation.update')"
                  :value="Operation.update"
                ></v-radio>
                <v-radio
                  :label="$t('keycloak.operation.create')"
                  :value="Operation.create"
                ></v-radio>
                <v-radio
                  :label="$t('keycloak.operation.delete')"
                  :value="Operation.delete"
                ></v-radio>
              </v-radio-group>

              <v-card-actions>
                <v-btn variant="flat" color="primary" @click="next">
                  {{ $t("keycloak.next") }}
                </v-btn>
                <v-btn @click="newUpload">
                  {{ $t("keycloak.previous") }}
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-stepper-window-item>

          <v-stepper-window-item value="3">
            <v-card :title="$t('keycloak.selectFormat')" flat>
              <v-radio-group v-model="format">
                <v-radio
                  :label="$t('keycloak.format.space')"
                  :value="Format.space"
                ></v-radio>
                <v-radio
                  :label="$t('keycloak.format.json')"
                  :value="Format.json"
                ></v-radio>
                <v-radio
                  :label="$t('keycloak.format.xml')"
                  :value="Format.xml"
                ></v-radio>
              </v-radio-group>

              <v-card-actions>
                <v-btn variant="flat" color="primary" @click="preview">
                  {{ $t("keycloak.next") }}
                </v-btn>
                <v-btn @click="prev">
                  {{ $t("keycloak.previous") }}
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-stepper-window-item>

          <v-stepper-window-item value="4">
            <v-card :title="$t('keycloak.preview')" flat>
              <v-card-actions>
                <v-btn variant="flat" color="primary" @click="process">
                  {{ $t("keycloak.confirm") }}
                </v-btn>
                <v-btn @click="cancelPreview">
                  {{ $t("keycloak.cancel") }}
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-stepper-window-item>

          <v-stepper-window-item value="5">
            <v-card :title="$t('keycloak.processing')" flat>
              <v-card-actions>
                <v-btn variant="flat" color="primary"> OK </v-btn>
              </v-card-actions>
            </v-card>
          </v-stepper-window-item>
        </v-stepper-window>
      </template>
    </v-stepper>
  </v-container>
</template>

<script lang="ts" setup>
import { ref } from "vue";

enum Operation {
  update,
  create,
  delete,
}

enum Format {
  space,
  json,
  xml,
}

const step = ref(0);

const operation = ref<Operation>(Operation.update);

const format = ref<Format>(Format.space);

function uploadFile() {
  step.value = step.value + 1;
}

function newUpload() {
  step.value = step.value - 1;
}

function prev() {
  step.value = step.value - 1;
}

function next() {
  step.value = step.value + 1;
}

function preview() {
  step.value = step.value + 1;
}

function cancelPreview() {
  step.value = step.value - 1;
}

function process() {
  step.value = step.value + 1;
}
</script>
