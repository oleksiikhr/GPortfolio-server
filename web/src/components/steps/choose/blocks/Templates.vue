<template>
  <div>
    <a-alert
      message="Choose the template you want to use"
      type="info"
      class="mb-20"
      show-icon
    />
    <div
      :class="['block_templates', {
        has_active: template
      }]"
    >
      <template-item
        v-for="(t, index) in templates"
        :key="index"
        :template="t"
        :class="{ active: t.name === template }"
        @click.native="onClick(t)"
      />
    </div>
  </div>
</template>

<script>
import TemplateItem from '@/components/steps/choose/templates/Item'
import templates from '@/data/templates'

export default {
  components: {
    TemplateItem
  },
  data() {
    return {
      templates
    }
  },
  computed: {
    template() {
      return this.$store.state.global.template
    }
  },
  methods: {
    onClick(template) {
      // Select and deselect
      this.$store.commit('global/UPDATE', {
        key: 'template',
        value: template.name === this.template ? '' : template.name
      })
    }
  }
}
</script>
