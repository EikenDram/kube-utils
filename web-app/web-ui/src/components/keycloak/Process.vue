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
                v-model="file"
                accept="*"
                :label="$t('keycloak.file')"
              ></v-file-input>

              <v-card-actions>
                <v-btn
                  variant="flat"
                  color="primary"
                  @click="uploadFile"
                  :disabled="disabledUpload"
                >
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
                  {{ $t("keycloak.preview") }}
                </v-btn>
                <v-btn @click="prev">
                  {{ $t("keycloak.previous") }}
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-stepper-window-item>

          <v-stepper-window-item value="4">
            <v-card :title="$t('keycloak.preview')" flat>
              <v-data-table
                :headers="previewData?.headers"
                :items="previewData?.items"
                :item-value="previewData?.key"
                :loading="loading"
                class="elevation-1"
              ></v-data-table>

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
              <div v-if="loading" class="text-center">
                <v-progress-circular indeterminate>
                  Loading...
                </v-progress-circular>
              </div>
              <v-sheet v-else> RESULT </v-sheet>

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
import { ref, computed } from "vue";
import { useAPI } from "@/plugins/axios";
import ResponseData from "@/types/responseData";
import { DataTableHeader } from "@/types/vuetify";

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

const api = useAPI();

const step = ref(0);

const uid = ref<string>();

const operation = ref<Operation>(Operation.update);

const format = ref<Format>(Format.space);

const loading = ref(true);

// upload file

const file = ref<File[] | undefined>();

const disabledUpload = computed(() => {
  let f = file.value ? file.value[0] : null;
  return !f;
});

function uploadFile() {
  let f = file.value ? file.value[0] : null;
  if (f)
    api.Keycloak.upload(f)
      .then((response: ResponseData) => {
        console.log(response.data);
        uid.value = response.data;
        step.value = step.value + 1;
      })
      .catch((err) => {
        console.log(err);
      });
}

function newUpload() {
  //2D: reset file
  step.value = step.value - 1;
}

function prev() {
  step.value = step.value - 1;
}

function next() {
  step.value = step.value + 1;
}

// preview file

type PreviewTable = {
  headers: DataTableHeader[];
  items: any[];
  key: string;
};

const previewData = ref<PreviewTable>();

function preview() {
  const data = {
    uid: uid.value,
    operation: Operation[operation.value],
    format: Format[format.value],
  };

  api.Keycloak.preview(data)
    .then((response: ResponseData) => {
      console.log(response.data);
      previewData.value = response.data;
      step.value = step.value + 1;
    })
    .catch((err) => {
      console.log(err);
    });
}

function cancelPreview() {
  step.value = step.value - 1;
}

// process file

function process() {
  const data = {
    uid: uid.value,
    operation: Operation[operation.value],
    format: Format[format.value],
  };

  api.Keycloak.process(data)
    .then((response: ResponseData) => {
      console.log(response.data);
      previewData.value = response.data;
      step.value = step.value + 1;
    })
    .catch((err) => {
      console.log(err);
    });
}
</script>
