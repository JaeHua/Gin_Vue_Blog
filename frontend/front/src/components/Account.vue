<template>
    <v-app>
      <v-navigation-drawer
        app
        permanent
        color="blue-grey"
        expand-on-hover
        rail
      >
        <v-list>
          <v-list-item>
            <v-list-item-avatar>
              <!-- <img v-if="!profileInfo.avatar" src="https://randomuser.me/api/portraits/men/85.jpg" alt="User" > -->
            <img  :src="userInfo?.avatar" alt="">

            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-title class="white--text">{{ userInfo?.name }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>

        <v-divider></v-divider>

        <v-list dense nav>
          <v-list-item @click="currentView = 'liked'">
            <v-list-item-icon>
              <v-icon>mdi-thumb-up</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>点赞文章</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-list-item @click="currentView = 'saved'">
            <v-list-item-icon>
              <v-icon>mdi-bookmark</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>收藏文章</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-list-item @click="currentView = 'settings'">
            <v-list-item-icon>
              <v-icon>mdi-account-cog</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>用户设置</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>

        <v-spacer></v-spacer>

        <v-list dense nav>

          <v-list-item @click="$router.push('/')">
            <v-list-item-icon>
              <v-icon>mdi-arrow-left-bold</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>返回博客</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-main>
        <v-container>
          <component :is="currentViewComponent"></component>
        </v-container>
      </v-main>
    </v-app>
  </template>

<script>
import UserSettings from '@/components/UserSettings.vue'
import LikedArticles from '@/components/LikedArticles.vue'
import SavedArticles from '@/components/SavedArticles.vue'

export default {
  name: 'YourComponentName',
  data () {
    return {
      currentView: 'settings'
    }
  },
  computed: {
    currentViewComponent () {
      switch (this.currentView) {
        case 'liked':
          return LikedArticles
        case 'saved':
          return SavedArticles
        case 'settings':
        default:
          return UserSettings
      }
    },
    userInfo () {
      return this.$store.state.userModule.userInfo
    }

  },
  created () {
  },
  methods: {

  }
}
</script>

  <style scoped>
  .v-list-item-title {
    color: white;
  }
  </style>
