<template>
  <div :class="`component_block cb_level_${level}`">
    <div class="cb_title">
      <a
        :id="id"
        :href="link"
        class="cb_link"
      >
        {{ text }}
      </a>
    </div>
    <div class="cb_content">
      <slot />
    </div>
  </div>
</template>

<script>
export default {
  name: 'Block',
  props: {
    title: {
      type: String,
      required: true
    },
    href: {
      type: String,
      default: ''
    },
    level: {
      type: [Number, String],
      default: 1,
      validator: (val) => +val > 0 && +val < 4
    }
  },
  computed: {
    id() {
      if (this.href) {
        return this.href[0] === '#' ? this.href.substring(1) : this.href
      }

      return this.title.trim().toLowerCase().replace(/[^a-z0-9]/g, '-')
    },
    link() {
      return '#' + this.id
    },
    text() {
      return (this.level === 1 ? '# ' : '') + this.title
    }
  }
}
</script>
