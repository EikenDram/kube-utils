<template>
  <v-container fluid>
    <v-data-table
      :headers="headers"
      :items="logs"
      item-value="uid"
      show-expand
      :loading="loading"
      class="elevation-1"
    >
      <template v-slot:expanded-row="{ columns, item }">
        <tr>
          <td :colspan="columns.length"><LogContent :log="item.value" /></td>
        </tr>
      </template>
    </v-data-table>
  </v-container>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { useAPI } from "@/plugins/axios";
import LogContent from "./LogContent.vue";
import { useI18n } from "vue-i18n";
import { DataTableHeader } from "@/types/vuetify";
import ResponseData from "@/types/responseData";

const { t } = useI18n(); // use as global scope

const api = useAPI();

const headers: DataTableHeader[] = [
  { key: "uid", title: t("keycloak.datetime") },
  { key: "stamp", title: t("keycloak.datetime") },
];

type Log = {
  uid: string;
  stamp: string;
};

const logs = ref<Log[]>([]);
const loading = ref(true);

function refresh() {
  api.Keycloak.getAll()
    .then((response: ResponseData) => {
      logs.value = response.data;
      console.log(response.data);
    })
    .catch((e: Error) => {
      console.log(e);
    })
    .finally(() => {
      loading.value = false;
    });
}

onMounted(() => {
  refresh();
});
</script>
