<template>
  <auto-step :menu="menu" />
</template>

<script>
import GithubWebsiteBlock from '@/components/steps/choose/websites/Github'
import TemplatesBlock from '@/components/steps/choose/Templates'
import WebsitesBlock from '@/components/steps/choose/Websites'
import StorageBlock from '@/components/steps/choose/Storage'
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
          component: WebsitesBlock,
          href: '#websites',
          title: 'Websites',
          children: [
            {
              component: GithubWebsiteBlock,
              href: '#website-github',
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
