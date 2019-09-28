<template>
  <step :menu="menu">
    <block title="Storage">
      <a-alert message="This functionality is in development" type="info" showIcon />
    </block>
    <block title="API">
      <a-icon type="github" />
      TODO
    </block>
    <block title="Templates">
      <templates-block />
    </block>
  </step>
</template>

<script>
import TemplatesBlock from '@/components/blocks/choose/Templates'
import Block from '@/components/Block'
import Step from '@/components/Step'

export default {
  components: {
    Step, Block, TemplatesBlock
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
