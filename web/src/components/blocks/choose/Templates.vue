<template>
  <block title="Templates">
    <a-alert
      message="Choose the template you want to use"
      type="info"
      class="mb-20"
      showIcon
    />
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
  </block>
</template>

<script>
import Block from '@/components/Block'

export default {
  components: {
    Block
  }
}
</script>
