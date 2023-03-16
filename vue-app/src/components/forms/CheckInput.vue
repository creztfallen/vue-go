<!-- eslint-disable vue/valid-template-root -->
<template>
  <div class="form-check">
    <input :checked="checked" :required="required" class="form-check-input" />
    <label :for="name" class="form-check-label">{{ label }}</label>
  </div>
</template>

<script>
export default {
  name: "CheckInput",
  props: ["label", "required", "name", "value", "checked"],
};
</script>