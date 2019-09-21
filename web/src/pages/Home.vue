<template>
  <div>
    <router-link :to="{ name: 'home' }">Home</router-link>
    <router-link :to="{ name: 'test' }">Test</router-link>
  </div>
</template>

<script>
export default {
  created() {
    console.log('HERE')
  }
}
</script>
