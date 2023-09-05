<template>
  <div>{{ props.log }}</div>
</template>

<script lang="ts" setup>
import { ref } from "vue";

const props = defineProps({
  log: String,
});

// need to do api call on mount
</script>
