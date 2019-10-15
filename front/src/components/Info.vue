<template>
<div class="info">
  <h1>Our TODOs:</h1>
  <p>
    List of user's todo items.
  </p>
  <fake-button type="All"></fake-button>
</div>
</template>

<script>
import FakeButton from "./FakeButton.vue"

export default {
  name: "Info",
  components: {
    FakeButton,
  },
}
</script>

<style>

</style>
