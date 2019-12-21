<template>
  <div>
    <p>Choose the template you want to use.</p>
    <div
      :class="['block_templates', 'mt-20', {
        has_active: activeTemplate
      }]"
    >
      <template-item
        v-for="(template, index) in templates"
        :key="index"
        :template="template"
        :class="{ active: template.name === activeTemplate }"
        @click.native="onClick(template)"
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
    activeTemplate() {
      return this.$store.state.global.template
    }
  },
  methods: {
    onClick(template) {
      const chooseTemplate = this.activeTemplate === template.name ? '' : template.name
      console.log('chooseTemplate', chooseTemplate)
      this.$store.commit('global/SET_TEMPLATE', chooseTemplate)
    }
  }
}
</script>
