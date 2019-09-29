<template>
  <div>
    <a-alert
      message="Choose the template you want to use"
      type="info"
      class="mb-20"
      showIcon
    />
    <div class="block_templates">
      <template-item
        v-for="(template, index) in templates"
        :key="index"
        :template="template"
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
  }
}
</script>
