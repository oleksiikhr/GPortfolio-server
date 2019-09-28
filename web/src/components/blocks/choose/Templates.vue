<template>
  <div class="block_templates">
    <a-card
      v-for="i in 10"
      :key="i"
      hoverable
      title="Some template"
    >
      <img
        :src="`https://picsum.photos/id/${i + 250}/200`"
        alt="example"
        slot="cover"
      />
      <a-badge
        slot="extra"
        status="error"
      />
    </a-card>
  </div>
</template>

<script>
export default {

}
</script>
