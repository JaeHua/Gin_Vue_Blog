<template >
<div>
    <v-card  class="mx-n15" min-width="300">
        <v-img :src="profileInfo.img">
            <v-card-title>
                <v-col align="center">
                    <v-avatar size="130" color="grey">
                        <img :src="profileInfo.avatar"  alt="">
                    </v-avatar>
                    <div class="mt-4 white--text">{{ profileInfo.name }}</div>
                </v-col>
            </v-card-title>
            <v-divider></v-divider>
        </v-img>
        <v-col>
            <div class="ma-3 text-h5 font-weight-bold">About Me:</div>
            <div class="align-center">
      <v-chip
        v-for="(item, index) in descArray"
        :key="index"
        :color="colors[index % colors.length]"
        class="mr-2"
        variant="flat"
      >
        {{ item }}
      </v-chip>
      </div>
        </v-col>
        <v-divider color="indigo"></v-divider>
        <v-list nav dense>
            <v-list-item class="mb-n3">
                <v-list-item-icon class="ma-3 ">
                    <v-icon  color="blue-darken-2">{{'mdi-qqchat'}}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>{{ profileInfo.qq }}</v-list-item-content>
            </v-list-item>
            <v-list-item class="mb-n3">
                <v-list-item-icon class="ma-3 ">
                    <v-icon  color="blue-lighten-2">{{'mdi-github'}}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                    <a href="https://github.com/JaeHua" target="_blank" style="color: inherit; text-decoration: none;">{{ profileInfo.github }}</a>
                </v-list-item-content>
            </v-list-item>
            <v-list-item class="mb-n3">
                <v-list-item-icon class="ma-3 ">
                    <v-icon  color="blue-lighten-2">{{'mdi-email'}}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                    {{ profileInfo.email }}
                </v-list-item-content>
            </v-list-item>
            <v-list-item class="mb-n3">
                <v-list-item-icon class="ma-3 ">
                    <v-icon  color="blue-lighten-2">{{'mdi-signal-variant'}}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                    {{ profileInfo.site }}
                </v-list-item-content>
            </v-list-item>
        </v-list>
    </v-card>
</div>
</template>
<script>

export default {
  data () {
    return {
      profileInfo: {
        id: 1
      },
      colors: ['blue', 'red', 'green']
    }
  },
  created () {
    this.getProfileInfo()
  },
  methods: {
    // 个人设置
    async getProfileInfo () {
      const { data: res } = await this.$http.get(`profile/${this.profileInfo.id}`)
      this.profileInfo = res.data
      console.log(this.profileInfo)
    }
  },
  computed: {
    descArray () {
      return String(this.profileInfo.desc).split('_')
    }
  }
}
</script>
