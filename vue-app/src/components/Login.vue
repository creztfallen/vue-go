<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr />
        <form-tag @myevent="submithandler" name="myform" event="myevent">
          <text-input v-model="email" label="Email" type="email" name="email" required="true">
          </text-input>
          <text-input v-model="password" label="Password" type="password" name="password" required="true">
          </text-input>
          <input type="submit" class="btn btn-primary" value="Login" />
        </form-tag>
        <hr />
      </div>
    </div>
  </div>
</template>

<script>
import FormTag from "./forms/FormTag.vue";
import TextInput from "./forms/TextInput.vue";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "login",
  components: {
    TextInput,
    FormTag,
  },
  data() {
    return {
      email: "",
      password: "",
    }
  },
  methods: {
    submitHandler() {
      console.log("submitHandler called - success!");
    }
  },

};
</script>