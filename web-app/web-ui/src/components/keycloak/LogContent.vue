<template>
  <div>
    <div v-if="loading" class="text-center">
      <v-progress-circular indeterminate>Loading...</v-progress-circular>
    </div>
    <div v-else>{{ data }}</div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { useAPI } from "@/plugins/axios";
import ResponseData from "@/types/responseData";

const api = useAPI();

const loading = ref(true);

const data = ref({});

const props = defineProps({
  log: String,
});

onMounted(() => {
  api.Keycloak.get(props.log ?? "")
    .then((response: ResponseData) => {
      data.value = response.data;
      console.log(response.data);
    })
    .catch((e: Error) => {
      console.log(e);
    })
    .finally(() => {
      loading.value = false;
    });
});

// need to do api call on mount
</script>
