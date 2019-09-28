<template>
  <div class="page home_page">
    <app-logo />
    <header class="mw_4k mc hp_header">
      <div class="hp_header__wrap">
        <div class="hp_header_left">
          <h1>Create Your Portfolio Online</h1>
          <p>Upload your data from social networks or fill in manually</p>
          <div class="ant-row-flex hp_header_left__actions">
            <a-button
              type="primary"
              size="large"
              @click="onClickTryIt"
            >
              Try it Free
            </a-button>
            <a-button
              size="large"
              icon="github"
              @click="onClickGithub"
            >
              Github
            </a-button>
          </div>
          <contributors />
        </div>
        <div class="hp_header_right">
          <hire-svg />
        </div>
      </div>
    </header>
    <wave-svg class="hp_wave" />
  </div>
</template>

<script>
import Contributors from '@/components/Contributors'
import { GENERATE_PAGE } from '@/router/routes'
import AppLogo from '@/components/global/Logo'
import { ORGANIZATION } from '@/scripts/links'
import HireSvg from '@/components/svg/Hire'
import WaveSvg from '@/components/svg/Wave'

export default {
  components: {
    AppLogo, HireSvg, WaveSvg, Contributors
  },
  methods: {
    onClickGithub() {
      window.open(ORGANIZATION, '_blank')
    },
    onClickTryIt() {
      this.$router.push({ name: GENERATE_PAGE })
    }
  }
}
</script>
