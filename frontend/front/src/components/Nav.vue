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

        <!-- Desc Section -->
        <div class="headerinfo">
          <p class="desc">{{ profileInfo.desc }}</p>
        </div>

        <!-- Icon Section -->
        <div class="icon-container">
          <div class="circle">
            <v-icon class="arrow">{{'mdi-arrow-left'}}</v-icon>
          </div>
          <div class="circle" @click="navigateTo('qq')">
            <v-icon class="qqchat">{{ 'mdi-qqchat' }}</v-icon>
          </div>
          <div class="circle" @click="navigateTo('github')">
            <v-icon class="github">{{ 'mdi-github' }}</v-icon>
          </div>
          <div class="circle" @click="navigateTo('email')">
            <v-icon class="email">{{ 'mdi-email' }}</v-icon>
          </div>
          <div class="circle" @click="navigateTo('signal')">
            <v-icon class="signal">{{ 'mdi-signal-variant' }}</v-icon>
          </div>
          <div class="circle">
            <v-icon class="arrow">{{ 'mdi-arrow-right' }}</v-icon>
          </div>
        </div>
      </v-img>
    </v-card>
  </div>
</template>
<script>
import { mapState, mapGetters } from 'vuex'

export default {
  data () {
    return {
      profileInfo: {
        email: 'admin'
      },
      colors: ['blue', 'red', 'green'],
      hash: {
        'qq.com': 'http://mail.qq.com',
        'gmail.com': 'http://mail.google.com',
        'sina.com': 'http://mail.sina.com.cn',
        '163.com': 'http://mail.163.com',
        '126.com': 'http://mail.126.com',
        'yeah.net': 'http://www.yeah.net/',
        'sohu.com': 'http://mail.sohu.com/',
        'tom.com': 'http://mail.tom.com/',
        'sogou.com': 'http://mail.sogou.com/',
        '139.com': 'http://mail.10086.cn/',
        'hotmail.com': 'http://www.hotmail.com',
        'live.com': 'http://login.live.com/',
        'live.cn': 'http://login.live.cn/',
        'live.com.cn': 'http://login.live.com.cn',
        '189.com': 'http://webmail16.189.cn/webmail/',
        'yahoo.com.cn': 'http://mail.cn.yahoo.com/',
        'yahoo.cn': 'http://mail.cn.yahoo.com/',
        'eyou.com': 'http://www.eyou.com/',
        '21cn.com': 'http://mail.21cn.com/',
        '188.com': 'http://www.188.com/',
        'foxmail.com': 'http://www.foxmail.com',
        'outlook.com': 'http://www.outlook.com'
      },
      eemail: ''
    }
  },
  computed: {
    ...mapState(['ID_email']),
    ...mapGetters(['isLoggedIn']),
    descArray () {
      return String(this.profileInfo.desc).split('_')
    }
  },
  created () {
    this.getAuthProfileInfo()
  },
  methods: {
    // 个人设置
    async getAuthProfileInfo () {
      const { data: res } = await this.$http.get(`profile/${this.profileInfo.email}`)
      this.profileInfo = res.data
      // console.log(this.profileInfo)
    },
    navigateTo (target) {
      switch (target) {
        case 'qq':
          window.location.href = `https://user.qzone.qq.com/${this.profileInfo.qqchat}`
          break
        case 'github':
          window.location.href = `https://github.com/${this.profileInfo.github}`
          break
        case 'signal':
          window.location.href = this.profileInfo.site
          break
        case 'email':
          this.eemail = this.profileInfo.email.split('@')[1] // 获取邮箱域
          window.location.href = `${this.hash[this.eemail]}/${this.profileInfo.email.split('@')[0]}`
          break
        default:
          break
      }
    }
  }

}
</script>

<style scoped>
/* Header Info */
.headerinfo {
  background-color: rgba(0, 0, 0, 0.5); /* 透明黑色 */
  color: #eaeadf;
  padding: 15px;
  border-radius: 5px;
  text-align: center;
  margin: auto;
  width: 60%; /* 调整宽度以适应小屏 */
  letter-spacing: 0;
  line-height: 30px;
}

/* Desc Text */
.desc {
  font-family: 'Pacifico', cursive;
  font-size: 24px;
  color: #f7f4f4;
  text-align: center;
  word-break: break-word; /* 强制长单词换行 */
  white-space: pre-wrap; /* 保留换行符并允许多行显示 */
  text-shadow: 2px 2px 6px rgba(255, 0, 0, 0.7),
               -2px -2px 6px rgba(0, 255, 0, 0.7),
               2px -2px 6px rgba(0, 0, 255, 0.7),
               -2px 2px 6px rgba(255, 255, 0, 0.7);
}

/* Icon Container */
.icon-container {
  display: flex;
  justify-content: center; /* 确保图标容器居中 */
  flex-wrap: wrap; /* 允许图标在小屏幕时换行 */
  gap: 10px; /* 图标之间的间距 */
  margin-top: 15px;
}

/* Circle */
.circle {
  background-color: rgba(0, 0, 0, 0.5); /* 透明黑色背景 */
  padding: 10px;
  border-radius: 50%; /* 圆形 */
  display: inline-flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.3s ease; /* 添加过渡效果 */
}

.circle:hover {
  transform: scale(1.2); /* 鼠标悬停时放大 */
}

/* 图标颜色 */
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

/* 媒体查询: 小屏幕适配 */
@media (max-width: 600px) {
  .desc {
    font-size: 18px; /* 调整desc大小 */
  }

  .icon-container {
    gap: 5px; /* 小屏幕时缩小图标间距 */
    margin-top: 10px;
  }

  .circle {
    padding: 8px; /* 小屏幕时缩小图标大小 */
  }

  .headerinfo {
    width: 90%; /* 小屏幕时扩大headerinfo宽度 */
  }
}

/* Avatar Hover */
.avatar-hover:hover img {
  animation: rotateAndZoom 0.5s forwards;
}

@keyframes rotateAndZoom {
  from {
    transform: rotate(0deg) scale(1);
  }
  to {
    transform: rotate(180deg) scale(1.2);
  }
}

/* 颜色渐变动画 */
@keyframes colorChange {
  0% { color: #f7f4f4; }
  25% { color: #ff99c8; }
  50% { color: #fcf6bd; }
  75% { color: #d0f4de; }
  100% { color: #a9def9; }
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

.github:hover,
.qqchat:hover,
.email:hover,
.signal:hover,
.arrow:hover {
  animation: zoomIn 0.3s forwards;
}
</style>
