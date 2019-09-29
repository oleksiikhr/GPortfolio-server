<template>
  <a-card
    title="Default"
    hoverable
  >
    <img
      slot="cover"
      :src="template.image"
      :alt="`Template: ${template.name}`"
      loading="lazy"
    >
    <a-avatar
      slot="extra"
      :alt="contributor.name"
      :src="contributor.image"
      :title="contributor.name"
      size="small"
    />
  </a-card>
</template>

<script>
import contributors from '@/data/contributors'

export default {
  props: {
    template: {
      type: Object,
      required: true
    }
  },
  computed: {
    contributor() {
      if (!this.template.author) {
        return {}
      }

      return contributors[this.template.author]
    }
  }
}
</script>
