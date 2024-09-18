<template>
  <v-card class="fill-height d-flex flex-column" flat>
    <v-img
      :src="profile.img"
      gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.7)"
      height="100%"
      class="imgs"
    >
      <v-container fluid fill-height class="pa-0">
        <v-row no-gutters align="center" justify="center" class="fill-height">
          <v-col cols="12" class="d-flex flex-column justify-center align-center text-center">
            <v-avatar size="200" class="mb-8 elevation-10">
              <v-img :src="profile.avatar" v-if="profile.avatar"></v-img>
              <v-icon size="120" color="grey lighten-2" v-else>mdi-account</v-icon>
            </v-avatar>

            <h1 class="text-h2 font-weight-bold white--text mb-4">{{ profile.name }}</h1>
            <p class="text-h5 white--text mb-4" style="max-width: 600px; min-height: 3em;">
              <TypewriterEffect :text="profile.desc" :speed="50" />
            </p>
            <!-- <p class="text-h5 white--text mb-6" style="max-width: 600px;">{{ profile.description }}</p> -->
            <p class="text-h6 white--text mb-6">
              <v-icon color="white" left>mdi-email</v-icon>
              {{ profile.email }}
            </p>
            <v-row justify="center" align="center" class="mb-6">
              <v-col cols="auto">
                <v-btn
                  color="white"
                  text
                  @click="goToGithub"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="px-2"
                >
                  <v-icon left>mdi-github</v-icon>
                  GitHub
                </v-btn>
              </v-col>
              <v-col cols="auto">
                <v-btn
                  color="white"
                  text
                  @click="drawer = true"
                  class="px-2"
                >
                  <v-icon left>mdi-pencil</v-icon>
                  编辑资料
                </v-btn>
              </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-container>
    </v-img>

    <v-navigation-drawer
      v-model="drawer"
      fixed
      right
      temporary
      width="400"
    >
      <v-card flat height="100%">
        <v-card-title class="text-h6">
          编辑个人资料
          <v-spacer></v-spacer>
          <v-btn icon @click="drawer = false">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-card-title>

        <v-card-text>
          <v-container>
            <v-row justify="center" class="mb-6">
              <v-avatar size="150" color="grey lighten-2">
                <v-img :src="editedProfile.avatar" v-if="editedProfile.avatar"></v-img>
                <v-icon size="80" v-else>mdi-account</v-icon>
              </v-avatar>
            </v-row>
            <v-row justify="center" class="mb-6">
              <v-btn color="primary" @click="triggerUpload">
                <v-icon left>mdi-camera</v-icon>
                更换头像
              </v-btn>
              <input
                type="file"
                ref="fileInput"
                accept="image/*"
                style="display: none"
                @change="onFileSelected"
              />
            </v-row>
            <v-text-field
              v-model="editedProfile.name"
              label="用户名"
              prepend-icon="mdi-account"
              :rules="[v => !!v || '用户名不能为空']"
              required
            ></v-text-field>
            <v-text-field
              v-model="editedProfile.desc"
              label="个人描述"
              prepend-icon="mdi-text"
              rows="3"
              auto-grow
            ></v-text-field>
            <v-row align="center" class="mb-6">
              <v-col cols="12" sm="8">
                <v-file-input
                  v-model="backgroundFile"
                  label="上传背景图片"
                  prepend-icon="mdi-image"
                  accept="image/*"
                  @change="onBackgroundSelected"
                ></v-file-input>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>

        <v-card-actions class="justify-center">
          <v-btn color="primary" @click="saveProfile" :loading="saving">
            保存修改
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-navigation-drawer>
  </v-card>
</template>
<script>
import TypewriterEffect from './TypewriterEffect.vue'
export default {
  components: {
    TypewriterEffect
  },
  data () {
    return {
      drawer: false,
      editedProfile: {},
      saving: false,
      backgroundFile: null
    }
  },
  computed: {
    profile () {
      return this.$store.state.userModule.userInfo
    }
  },
  created () {
    this.editedProfile = { ...this.profile }
  },
  methods: {
    triggerUpload () {
      this.$refs.fileInput.click()
    },
    onFileSelected (event) {
      const file = event.target.files[0]
      if (file) {
        const reader = new FileReader()
        reader.onload = (e) => {
          this.editedProfile.avatar = e.target.result
          this.uploadImage(file, 'avatar')
        }
        reader.readAsDataURL(file)
      }
    },
    onBackgroundSelected (file) {
      if (file) {
        this.uploadImage(file, 'background')
      }
    },
    uploadImage (file, type) {
      const formData = new FormData()
      formData.append('file', file)

      this.$http.post('upload', formData, {
        headers: this.headers
      }).then(response => {
        const url = response.data.url
        if (type === 'avatar') {
          this.editedProfile.avatar = url
        } else if (type === 'background') {
          this.editedProfile.img = url
        }
        this.$toast.success(`${type === 'avatar' ? '头像' : '背景'}上传成功`, { timeout: 3000 })
      }).catch(() => {
        this.$toast.error(`${type === 'avatar' ? '头像' : '背景'}上传失败`, { timeout: 3000 })
      })
    },
    saveProfile () {
      this.saving = true
      this.$http.put('editProfile', this.editedProfile)
        .then(() => {
          this.$store.commit('userModule/SET_USERINFO', this.editedProfile)
          this.saving = false
          this.drawer = false
          this.$toast.success('个人资料已成功更新！', { timeout: 3000 })
        }).catch(() => {
          this.saving = false
          this.$toast.error('更新失败，请重试', { timeout: 3000 })
        })
    },
    goToGithub () {
      window.location.href = 'https://github.com/JaeHua'
    }
  }
}
</script>
<style scoped>
.v-avatar {
  transition: all 0.3s ease-in-out;
}
.v-avatar:hover {
  transform: scale(1.05);
  box-shadow: 0 8px 16px rgba(0,0,0,0.2);
}
.v-btn {
  transition: all 0.3s ease-in-out;
}
.v-btn:hover {
  background-color: rgba(255, 255, 255, 0.1) !important;
}
.fill-height {
  height: 100%;
}
.flex-grow-1 {
  flex-grow: 1;
}
.imgs {
  min-height: 674px;
  max-height: 675px;
}
</style>
