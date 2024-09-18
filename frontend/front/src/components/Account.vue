<template>
  <v-app>
    <!-- Top App Bar with Hamburger Icon for Mobile, only visible on small screens -->
    <v-app-bar v-if="isMobile" app color="#1976d2" dense>
      <v-app-bar-nav-icon @click="drawer = !drawer" class="white--text"></v-app-bar-nav-icon>
      <v-toolbar-title class="white--text">List</v-toolbar-title>
    </v-app-bar>

    <!-- Navigation Drawer -->
    <v-navigation-drawer
      v-model="drawer"
      app
      :permanent="!isMobile"
      :width="isMobile ? 240 : 300"
      color="#1976d2"
    >
      <v-list>
        <v-list-item class="px-2 py-4">
          <v-list-item-avatar size="48">
            <v-img :src="userInfo?.avatar" alt="User Avatar" />
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title class="text-h6 white--text">{{ userInfo?.name }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <v-divider class="mx-3 mb-2" color="#5eaec0"></v-divider>

      <v-list dense nav>
        <v-list-item
          v-for="item in menuItems"
          :key="item.title"
          @click="currentView = item.view"
          :class="{ 'v-item--active': currentView === item.view }"
          class="mb-2"
        >
          <v-list-item-icon>
            <v-icon :color="currentView === item.view ? 'white' : '#5eaec0'">{{ item.icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title :class="{ 'white--text': currentView === item.view, 'text-h7': true }">
              {{ item.title }}
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <template v-slot:append>
        <v-divider class="mx-3 mt-2" color="#5eaec0"></v-divider>
        <v-list dense nav>
          <v-list-item @click="$router.push('/')" class="mt-2">
            <v-list-item-icon>
              <v-icon color="#5eaec0">mdi-arrow-left-bold</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title class="text-h7" style="color: #5eaec0;">返回博客</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </template>
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
  components: { UserSettings, LikedArticles, SavedArticles },

  data () {
    return {
      drawer: true, // 控制抽屉的显示与隐藏
      currentView: 'settings',
      menuItems: [
        { title: '点赞文章', icon: 'mdi-thumb-up', view: 'liked' },
        { title: '收藏文章', icon: 'mdi-bookmark', view: 'saved' },
        { title: '用户设置', icon: 'mdi-account-cog', view: 'settings' }
      ]
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
    },
    isMobile () {
      return this.$vuetify.breakpoint.smAndDown
    }
  },
  watch: {
    isMobile (val) {
      // 当切换到移动设备时，自动关闭抽屉
      if (val) this.drawer = false
    }
  }
}
</script>

<style scoped>
.v-list-item {
  border-radius: 0 25px 25px 0;
  margin-right: 12px;
}

.v-item--active {
  background-color: #4998a5 !important;
}

.v-list-item:hover:not(.v-item--active) {
  background-color: rgba(94, 174, 192, 0.1) !important;
}

.v-navigation-drawer ::v-deep .v-navigation-drawer__border {
  display: none;
}

.v-main {
  background-color: white; /* 确保主内容背景色不变暗 */
}

@media (max-width: 600px) {
  .v-navigation-drawer {
    width: 240px;
  }
}

.container {
  width: 100%;
  padding: 0px;
  margin-right: auto;
  margin-left: auto;
}
</style>
