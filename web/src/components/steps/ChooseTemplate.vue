<template>
  <step :menu="menu">
    <h2 id="storage" class="title"># Storage</h2>
    <a-alert message="This functionality is in development" type="info" showIcon />
    <h2 id="api" class="title"># API</h2>
    <a-icon type="github" />
    <h2 id="templates" class="title"># Templates</h2>
    <a-card
      hoverable
      style="width: 240px"
    >
      <img
        alt="example"
        src="https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png"
        slot="cover"
      />
      <a-card-meta
        title="Europe Street beat">
        <template slot="description">www.instagram.com</template>
      </a-card-meta>
    </a-card>
  </step>
</template>

<script>
import Step from '@/components/Step'

export default {
  components: {
    Step
  },
  data() {
    return {
      menu: [
        { href: '#storage', title: 'Storage' },
        { href: '#api', title: 'API' },
        { href: '#templates', title: 'Templates' },
      ],
      storage: [
        { title: 'Ant Design Title 1' }
      ]
    }
  }
}
</script>
