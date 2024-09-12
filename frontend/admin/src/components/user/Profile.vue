<template>
    <div>
        <a-card>
            <h2>个人设置</h2>
        <a-form-model :model="profileInfo">
        <a-form-model-item label="作者名称" >
            <a-input style="width: 200px;" v-model="profileInfo.name"  :value="profileInfo.name"></a-input>
        </a-form-model-item>

        <a-form-model-item label="作者描述" >
            <a-input type="textarea" style="width: 40%;" v-model="profileInfo.desc" :value="profileInfo.desc"></a-input>
        </a-form-model-item>

        <a-form-model-item label="QQ号" >
            <a-input style="width: 200px;" v-model="profileInfo.qq" :value="profileInfo.qq"></a-input>
        </a-form-model-item>

        <!-- <a-form-model-item label="邮箱" >
            <a-input style="width: 200px;" v-model="profileInfo.email" :value="profileInfo.email"></a-input>
        </a-form-model-item> -->

        <a-form-model-item label="技术网站" >
            <a-input style="width: 300px;"  v-model="profileInfo.site" :value="profileInfo.site"></a-input>
        </a-form-model-item>

        <a-form-model-item label="头像">
            <a-upload
            name="file"
            :multiple="false"
            :action="upUrl"
            :headers="headers"
            @change="avatarChange"
            listType="picture"
          >
            <a-button> <a-icon type="upload" /> 点击上传 </a-button>
          </a-upload>
          <br />
          <template v-if="profileInfo.avatar">
            <img :src="profileInfo.avatar" style="height: 100px;width: 100px;">
          </template>
        </a-form-model-item>

        <a-form-model-item label="头像背景图" >
            <a-upload
            name="file"
            :multiple="false"
            :action="upUrl"
            :headers="headers"
            @change="imgChange"
            listType="picture"
          >
            <a-button> <a-icon type="upload" /> 点击上传 </a-button>
          </a-upload>
          <br />
          <template v-if="profileInfo.img">
            <img :src="profileInfo.img" style="height: 100px;width: 100px;">
          </template>
        </a-form-model-item>
        <a-form-model-item label="Github" >
            <a-input style="width: 200px;" v-model="profileInfo.github" :value="profileInfo.github"></a-input>
        </a-form-model-item>

        <a-form-model-item>
            <a-button type="danger" style="margin-right: 15px;" @click="updateProfile">更新</a-button>
        </a-form-model-item>
        </a-form-model>
    </a-card>
    </div>
</template>

<script>
import URL from '@/plugin/axios'
export default {
  data () {
    return {
      profileInfo: {
        name: '',
        qq: '',
        desc: '',
        email: 'jaelele@163.com',
        site: '',
        img: '',
        github: '',
        avatar: ''
      },
      upUrl: URL + '/upload',
      headers: {}
    }
  },
  created () {
    this.getProfileInfo()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
  },
  methods: {
    // 获取用户信息
    async getProfileInfo () {
      const { data: res } = await this.$http.get(`profile/${this.profileInfo.email}`)
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      if (!res.data) {
        return this.$message.error('No data returned from API.')
      }

      this.profileInfo = res.data
    },
    // 上传头像
    avatarChange (info) {
      if (info.file.status !== 'uploading') {
        // console.log(info.file, info.fileList)
      }
      if (info.file.status === 'done') {
        this.$message.success(`${info.file.name} 上传成功`)
        const imgUrl = info.file.response.url
        this.profileInfo.avatar = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error(`${info.file.name} 上传失败.`)
      }
    },
    // 上传头像背景图
    imgChange (info) {
      if (info.file.status !== 'uploading') {
        // console.log(info.file, info.fileList)
      }
      if (info.file.status === 'done') {
        this.$message.success(`${info.file.name} 上传成功`)
        const imgUrl = info.file.response.url
        this.profileInfo.img = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error(`${info.file.name} 上传失败.`)
      }
    },
    async updateProfile () {
      console.log(this.profileInfo)
      const { data: res } = await this.$http.put(`profile/${this.profileInfo.id}`, this.profileInfo)
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.$message.success('个人信息修改成功')
      this.$router.push('/index')
    }
  }
}
</script>
