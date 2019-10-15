<template>
  <div>
    <button class="button">
      {{ type }}
    </button>
  </div>
</template>

<script>
  export default {
    name: "FakeButton",
    props: {
      type: {
        type: String,
        required: false,
        default: "all"
      },
    },
  }
</script>

<style scoped lang="scss">
  .button {
    margin-top: 50px;
    background-color: white;
    border: none;
    font-size: 14px;
    font-family: 'Fira Code', sans-serif;
    padding: 10px 19px;
    font-weight: 600;
    display: block;
  }
</style>
