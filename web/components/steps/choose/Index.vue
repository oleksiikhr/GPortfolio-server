<template>
  <auto-step :menu="menu" />
</template>

<script>
import GithubApiBlock from '@/components/steps/choose/api/Github'
import TemplatesBlock from '@/components/steps/choose/Templates'
import StorageBlock from '@/components/steps/choose/Storage'
import ApiBlock from '@/components/steps/choose/Api'
import AutoStep from '@/components/steps/AutoStep'

export default {
  components: {
    AutoStep
  },
  data() {
    return {
      menu: [
        {
          component: StorageBlock,
          href: '#storage',
          title: 'Storage'
        },
        {
          component: ApiBlock,
          href: '#api',
          title: 'API',
          children: [
            {
              component: GithubApiBlock,
              href: '#api-github',
              title: 'Github'
            }
          ]
        },
        {
          component: TemplatesBlock,
          href: '#template',
          title: 'Templates'
        }
      ]
    }
  }
}
</script>
