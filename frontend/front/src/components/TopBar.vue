<template>
  <div class="navv">
    <v-app-bar app color="primary" flat>
      <!-- logo -->
      <v-avatar class="mx-12" size="40" color="white">
        <v-img @click="gobackend" src="../assets/logo.png" alt=""></v-img>
      </v-avatar>

      <v-container class="py-0 fill-height">
        <v-btn @click="$router.push('/')" text color="white">首页</v-btn>
        <v-btn v-for="item in cateList" :key="item.id" text color="white" @click="$router.push(`/category/${item.id}`)">
          {{ item.name }}
        </v-btn>
      </v-container>
      <v-spacer></v-spacer>

      <v-responsive max-width color="white">
        <v-text-field
          ref="searchField"
          dense
          flat
          hide-details
          rounded
          dark
          solo-inverted
          append-icon="mdi-magnify"
          @click:append="goToResult"
          @keyup.enter="goToResult"
          v-model="searchQuery"
          placeholder="Search..."
        ></v-text-field>
      </v-responsive>

      <v-menu offset-y>
        <template v-slot:activator="{ on, attrs }">
          <v-avatar v-bind="attrs" v-on="on" class="ml-4" size="40" color="white">
            <v-icon v-if="!profileInfo.avatar">mdi-account-circle</v-icon>
            <img v-else :src="profileInfo.avatar" alt="">

          </v-avatar>
        </template>
        <v-list>
          <v-list-item v-if="!isLoggedIn" @click="loginDialog = true">
            <v-list-item-title>登录</v-list-item-title>
          </v-list-item>
          <v-list-item v-if="!isLoggedIn" @click="registerDialog = true">
            <v-list-item-title>注册</v-list-item-title>
          </v-list-item>
          <v-list-item v-if="isLoggedIn" @click="$router.push('/account')">
            <v-list-item-title>我的空间</v-list-item-title>
          </v-list-item>
          <v-list-item v-if="isLoggedIn" @click="logout">
            <v-list-item-title>退出登录</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>

      <!-- Login Dialog -->
      <v-dialog v-model="loginDialog" max-width="400">
        <v-card>
          <v-card-title class="headline">登录</v-card-title>
          <v-card-text>
            <v-text-field
              label="邮箱"
              v-model="userL.email"
              prepend-icon="mdi-account-tie"
              :rules="[rules.required, rules.maxLength]"
            ></v-text-field>
            <v-text-field
              label="密码"
              v-model="userL.password"
              prepend-icon="mdi-lock"
              type="password"
              :rules="[rules.required, rules.maxLength]"
            ></v-text-field>
            <v-btn block color="primary" @click="login">登录</v-btn>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn text @click="clearForm()">关闭</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>

      <!-- Register Dialog -->
      <v-dialog v-model="registerDialog" max-width="400">
        <v-card>
          <v-card-title class="headline">注册</v-card-title>
          <v-card-text>
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-text-field
                label="用户名"
                v-model="user.username"
                prepend-icon="mdi-account-tie"
                :rules="[rules.required, rules.maxLength]"
              ></v-text-field>
              <v-text-field
                label="邮箱"
                v-model="user.email"
                prepend-icon="mdi-email"
                type="email"
                :rules="[rules.required]"
              ></v-text-field>
              <v-text-field
                label="密码"
                v-model="user.password"
                prepend-icon="mdi-lock"
                type="password"
                :rules="[rules.required, rules.maxLength]"
              ></v-text-field>
              <v-text-field
                label="确认密码"
                v-model="confirmPassword"
                prepend-icon="mdi-lock"
                type="password"
                :rules="[rules.required, rules.maxLength, rules.matchPassword]"
              ></v-text-field>
              <v-row>
                <v-col>
                  <v-text-field
                    label="验证码"
                    v-model="verificationCode"
                    prepend-icon="mdi-shield-check"
                  ></v-text-field>
                </v-col>
                <v-col cols="auto">
                  <v-btn :disabled="timer > 0 || !valid" @click="sendVerificationCode">
                    {{ timer > 0 ? `重新发送 (${timer})` : '发送验证码' }}
                  </v-btn>
                </v-col>
              </v-row>
              <v-btn block color="secondary" @click="register" :disabled="!valid">注册</v-btn>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn text @click="clearForm()">关闭</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-app-bar>
  </div>
</template>

<script>
import { mapState, mapActions, mapGetters } from 'vuex'

export default {
  data () {
    return {
      cateList: [],
      searchQuery: '',
      loginDialog: false,
      registerDialog: false,
      confirmPassword: '',
      verificationCode: '',
      user: {
        username: '',
        email: '',
        password: '',
        role: 2
      },
      userL: {
        email: '',
        password: '',
        role: 2
      },
      valid: false,
      timer: 0,
      rules: {
        required: value => !!value || '不能为空',
        maxLength: value => value.length <= 20 || '不能超过20个字符',
        matchPassword: value => value === this.user.password || '两次密码必须一致'
      },
      profileInfo: {}
    }
  },
  computed: {
    ...mapState(['ID_email']),
    ...mapGetters(['isLoggedIn', 'getProfile'])
  },
  created () {
    this.GetCateList()
    const token = window.sessionStorage.getItem('token')
    const email = window.sessionStorage.getItem('email')
    if (token) {
      // 刷新页面会改变vuex状态，需要mock
      this._login(email)
    }
    // console.log(this.ID_email)
    this.getProfileInfo()
    // console.log(this.profileInfo)
  },
  methods: {
    ...mapActions(['_login', '_logout']),
    async getProfileInfo () {
      const { data: res } = await this.$http.get(`profile/${this.ID_email}`)
      this.profileInfo = res.data
    },
    async sendVerificationCode () {
      this.startTimer()
      const Qemial = { mail: this.user.email }
      const { data: res } = await this.$http.post('register/getcode', Qemial)
      console.log(res)
      return this.$toast.success('验证码发送成功', {
        timeout: 3000
      })
    },

    startTimer () {
      this.timer = 60
      const interval = setInterval(() => {
        this.timer--
        if (this.timer <= 0) {
          clearInterval(interval)
        }
      }, 1000)
    },
    async GetCateList () {
      const { data: res } = await this.$http.get('category')
      this.cateList = res.data
    },
    gobackend () {
      window.location.href = 'http://localhost:3000/admin'
    },
    goToResult () {
      const query = this.searchQuery.trim()
      if (query) {
        const currentQuery = this.$route.query.q
        if (currentQuery !== query) {
          this.$router.push({ path: '/result', query: { q: query } })
          this.searchQuery = '' // 清空输入框
        }
      }
    },
    clearForm () {
      this.registerDialog = false
      this.user.username = ''
      this.user.password = ''
      this.user.email = ''
      this.userL.email = ''
      this.confirmPassword = ''
      this.loginDialog = false
      this.userL.username = ''
      this.userL.password = ''
    },
    async register () {
      if (this.$refs.form.validate()) {
        const { data: codeRes } = await this.$http.post('register/verify', {
          mail: this.user.email,
          vcode: this.verificationCode
        })

        if (codeRes.status !== 200) {
          return this.$toast.error(codeRes.message, { timeout: 3000 })
        }

        const { data: res } = await this.$http.post('user/add', this.user)
        if (res.status !== 200) {
          this.registerDialog = false
          return this.$toast.error(res.message, { timeout: 3000 })
        }

        this.registerDialog = false
        return this.$toast.success('注册成功', { timeout: 3000 })
      }
    },
    async login () {
      const { data: res } = await this.$http.post('userlogin', this.userL)
      if (res.status !== 200) {
        this.loginDialog = false
        return this.$toast.error(res.message, {
          timeout: 3000
        })
      }
      this.loginDialog = false
      window.sessionStorage.setItem('token', res.token)
      window.sessionStorage.setItem('email', this.userL.email)
      this._login(this.userL.email)
      this.getProfileInfo()

      return this.$toast.success('登陆成功', {
        timeout: 3000
      })
    },
    logout () {
      window.sessionStorage.removeItem('token')
      this.profileInfo = {}
      this._logout()
      this.$toast.success('已退出登录', {
        timeout: 3000
      })
    }
  }
}
</script>

<style scoped>
.navv {
  z-index: 1000;
}
</style>
