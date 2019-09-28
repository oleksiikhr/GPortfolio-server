<template>
  <div class="component_block">
    <div class="cb_title">
      <a
        :id="id"
        :href="link"
        class="ant-anchor"
      >
        # {{ title }}
      </a>
    </div>
    <div class="cb_content">
      <slot />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      required: true
    },
    level: { // for Step.vue
      type: Number,
      default: 1,
      validator: (val) => val > 0
    }
  },
  computed: {
    id() {
      return this.title.trim().toLowerCase()
    },
    link() {
      return '#' + this.id
    }
  }
}
</script>
