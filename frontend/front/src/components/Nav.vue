<template>
  <div>
    <v-card class="mx-0" size="300px">
      <v-img :src="profileInfo.img">
        <v-card-title>
          <v-col align="center">
            <div>
              <v-avatar size="150" color="grey" class="mt-12 avatar-hover">
                <img :src="profileInfo.avatar" alt="">
              </v-avatar>
            </div>
            <div class="mt-4 white--text">{{ profileInfo.name }}</div>
          </v-col>
        </v-card-title>
        <div class="headerinfo">
          <p class="desc">{{ profileInfo.desc }}</p>
        </div>

        <div class="icon-container">
          <div class="circle">
            <v-icon class="arrow">{{'mdi-arrow-left'}}</v-icon>
          </div>
          <div class="circle">
            <v-icon class="qqchat">{{ 'mdi-qqchat' }}</v-icon>
          </div>
          <div class="circle">
            <v-icon class="github">{{ 'mdi-github' }}</v-icon>
          </div>
          <div class="circle">
            <v-icon class="email">{{'mdi-email'}}</v-icon>
          </div>
          <div class="circle">
            <v-icon class="signal">{{'mdi-signal-variant'}}</v-icon>
          </div>
          <div class="circle">
            <v-icon class="arrow">{{'mdi-arrow-right'}}</v-icon>
          </div>
        </div>
      </v-img>

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

<style scoped>
.headerinfo {
  background-color: rgba(0, 0, 0, 0.5); /* 透明黑色 */
  color: #eaeadf;
  padding: 15px; /* 内边距 */
  border-radius: 5px; /* 圆角 */
  text-align: center; /* 文字居中 */
  margin: auto;
  width: 34%;
  height: 60px;
  letter-spacing: 0;
  line-height: 30px;
}

.icon-container {
  display: flex; /* 使用Flexbox布局 */
  gap: 10px; /* 图标之间的间距 */
  margin-left: 40%;
  margin-top: 15px;
}

.circle {
  background-color: rgba(0, 0, 0, 0.5); /* 透明黑色背景 */
  padding: 10px; /* 内边距 */
  border-radius: 50%; /* 圆形 */
  display: inline-flex; /* 使其成为内联块并使内容居中 */
  align-items: center; /* 垂直居中 */
  justify-content: center; /* 水平居中 */
}

.qqchat {
  color: #00acee; /* QQ聊天图标颜色 */
}

.github {
  color: #333; /* GitHub图标颜色 */
}

.email {
  color: yellow;
}

.signal {
  color: rgb(255, 255, 255);
}

.arrow {
  color: #1295da;
}

/* 添加旋转和放大效果 */
@keyframes rotateAndZoom {
  from {
    transform: rotate(0deg) scale(1);
  }
  to {
    transform: rotate(180deg) scale(1.2);
  }
}

.avatar-hover:hover img {
  animation: rotateAndZoom 0.5s forwards; /* 0.5秒的动画，向前播放 */
}

/* 添加图标放大效果 */
@keyframes zoomIn {
  from {
    transform: scale(1);
  }
  to {
    transform: scale(1.2);
  }
}

 .github:hover {
  animation: zoomIn 0.3s forwards; /* 0.3秒的动画，向前播放 */
}
.qqchat:hover  {
  animation: zoomIn 0.3s forwards; /* 0.3秒的动画，向前播放 */
}
.email:hover  {
  animation: zoomIn 0.3s forwards; /* 0.3秒的动画，向前播放 */
}
.signal:hover  {
  animation: zoomIn 0.3s forwards; /* 0.3秒的动画，向前播放 */
}
.arrow:hover  {
  animation: zoomIn 0.3s forwards; /* 0.3秒的动画，向前播放 */
}
.desc {
      font-family: 'Pacifico', cursive;
      font-size: 24px;
      color: #f7f4f4;
      text-align: center;
      animation: fadeIn 2s ease-in-out, moveUpDown 3s infinite alternate, colorChange 5s infinite alternate;
      text-shadow: 2px 2px 6px rgba(255, 0, 0, 0.7),
                   -2px -2px 6px rgba(0, 255, 0, 0.7),
                   2px -2px 6px rgba(0, 0, 255, 0.7),
                   -2px 2px 6px rgba(255, 255, 0, 0.7);
    }

    @keyframes fadeIn {
      from { opacity: 0; }
      to { opacity: 1; }
    }

    @keyframes moveUpDown {
      from { transform: translateY(10px); }
      to { transform: translateY(-10px); }
    }

    @keyframes colorChange {
      0% { color: #f7f4f4; }
      25% { color: #ff99c8; }
      50% { color: #fcf6bd; }
      75% { color: #d0f4de; }
      100% { color: #a9def9; }
    }
</style>
