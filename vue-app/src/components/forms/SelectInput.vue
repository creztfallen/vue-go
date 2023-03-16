<!-- eslint-disable vue/valid-template-root -->
<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{ label }}</label>
    <select :name="name" :id="name" :required="required" class="form-select">
      <option v-for="option in items" :value="option.value" :key="option.value">
        {{ option.text }}
      </option>
    </select>
  </div>
</template>

<script>
export default {
  name: "SelectInput",
  props: ["items", "name", "required", "label"],
};
</script>