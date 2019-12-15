<template>
  <div>
    <a-button @click="onClick">Click</a-button>
    Github
  </div>
</template>

<script>
import { REGEX_PASS } from '@/data/auth'
import { popup } from '@/scripts/window'

export default {
  methods: {
    onClick() {
      // TODO Move logic outside
      const w = popup('/api/github/oauth/redirect', 'Github')

      w.addEventListener('beforeunload', (evt) => {
        const response = evt.target.body.innerText

        console.log('code', response)
        if (REGEX_PASS.test(response)) {
          const [login, password] = response.split('@')
          localStorage.setItem('api-github', JSON.stringify({ login, password }))
          console.log('Success!')
          // TODO Success, store to localStorage
        } else {
          console.log('Bad')
        }
      })
    }
  }
}
</script>
