<template>
  <v-container fluid>
    <v-data-table
      v-model.expanded="expanded"
      :headers="headers"
      :items="logs"
      item-value="uid"
      show-expand
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
import { ref } from "vue";
import LogContent from "./LogContent.vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n(); // use as global scope

type DataTableHeader = {
  key: string;
  title: string;
  colspan?: number;
  rowspan?: number;
  fixed?: boolean;
  align?: "start" | "end";
  width?: number;
  minWidth?: string;
  maxWidth?: string;
  sortable?: boolean;
};

const headers: DataTableHeader[] = [
  { key: "date", title: t("keycloak.datetime") },
];

const expanded = ref([]);

const logs = ref([
  {
    uid: "122131231",
    date: "datetime",
  },
  {
    uid: "122131232",
    date: "datetime",
  },
  {
    uid: "122131233",
    date: "datetime",
  },
  {
    uid: "122131234",
    date: "datetime",
  },
  {
    uid: "122131235",
    date: "datetime",
  },
]);
</script>
