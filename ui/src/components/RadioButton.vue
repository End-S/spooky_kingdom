<script setup lang="ts">
import isNil from "lodash-es/isNil";

const props = defineProps<{
  modelValue: string;
  nativeValue: string;
}>();

defineEmits<{
  (e: "update:modelValue", modelValue: string): void;
}>();

const isSelected = () => {
  return !isNil(props.modelValue) && props.modelValue === props.nativeValue;
};
</script>

<template>
  <o-button
    :v-model="modelValue"
    @click="$emit('update:modelValue', nativeValue)"
    :class="[
      { 'bg-cyan-200': isSelected(), 'dark:bg-cyan-800': isSelected() },
      'border',
      'border-black',
      'dark:border-white',
      '-mr-px',
      '!rounded-none',
    ]"
  >
    <slot></slot>
  </o-button>
</template>
