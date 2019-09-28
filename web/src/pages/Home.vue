<template>
  <div class="page home_page">
    <div class="mc tc">
      <logo-svg class="logo" />
    </div>
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
          <div class="hp_header_left__contributors">
            <a
              v-for="i in 50"
              :key="i"
              href="#"
            >
              <a-avatar
                icon="user"
                alt="TODO"
              />
            </a>
            <a-popover placement="bottom">
              <template slot="content">
                <span>Become a Contributor</span>
              </template>
              <a href="https://github.com/GPortfolio/GPortfolio" target="_blank">
                <a-avatar type="primary" style="backgroundColor: #3f3d56" icon="plus" alt />
              </a>
            </a-popover>
          </div>
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
import { REPOSITORY_GPORTFOLIO } from '../scripts/links'
import { GENERATE_PAGE } from '../router/routes'
import HireSvg from '@/components/svg/Hire'
import WaveSvg from '@/components/svg/Wave'
import LogoSvg from '@/components/svg/Logo'

export default {
  components: {
    HireSvg, WaveSvg, LogoSvg
  },
  methods: {
    onClickGithub() {
      window.open(REPOSITORY_GPORTFOLIO, '_blank')
    },
    onClickTryIt() {
      this.$router.push({ name: GENERATE_PAGE })
    }
  }
}
</script>
