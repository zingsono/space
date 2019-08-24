<template>
  <div>
    <h1>{{ msg }}</h1>
    <h2>h2</h2>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<style scoped>
  div{background-color: #0a6999;}
</style>