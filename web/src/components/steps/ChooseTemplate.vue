<template>
  <step :menu="menu">
    <storage-block />
    <api-block />
    <templates-block />
  </step>
</template>

<script>
import TemplatesBlock from '@/components/blocks/choose/Templates'
import StorageBlock from '@/components/blocks/choose/Storage'
import ApiBlock from '@/components/blocks/choose/Api'
import Step from '@/components/Step'

export default {
  components: {
    Step, ApiBlock, TemplatesBlock, StorageBlock
  },
  data() {
    return {
      menu: [
        { href: '#storage', title: 'Storage' },
        { href: '#api', title: 'API' },
        { href: '#templates', title: 'Templates' },
      ]
    }
  }
}
</script>
