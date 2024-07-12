<template>
<div>
    <a-card>
        <h3>{{id?'编辑文章':'新增文章'}}</h3>
        <a-form-model :model="artInfo" ref="artInfoRef" :rules="artInfoRules" :hideRequiredMark="true">
        <a-form-model-item label="文章标题" prop="title">
            <a-input style="width: 200px;" v-model="artInfo.title"></a-input>
        </a-form-model-item>
        <a-form-model-item label="文章分类" prop="cid">
            <a-select style="width: 200px;" @change="Catechange" v-model="artInfo.cid" placeholder="请选择分类">
                <a-select-option v-for="item in Catelist" :key="item.id" :value="item.id">{{item.name}}</a-select-option>
            </a-select>
        </a-form-model-item>
        <a-form-model-item label="文章描述" prop="desc">
            <a-input type="textarea" v-model="artInfo.desc"></a-input>
        </a-form-model-item>
        <a-form-model-item label="文章缩略图" prop="img">
            <a-upload
            name="file"
            :multiple="false"
            :action="upUrl"
            :headers="headers"
            @change="upChange"
            listType="picture"
          >
            <a-button> <a-icon type="upload" /> 点击上传 </a-button>
          </a-upload>
          <br />
          <template v-if="id">
            <img :src="artInfo.img" style="height: 100px;width: 100px;">
          </template>
        </a-form-model-item>
        <a-form-model-item label="文章内容" prop="content">
          <Editor v-model="artInfo.content"></Editor>
        </a-form-model-item>
        <a-form-model-item>
            <a-button type="danger" style="margin-right: 15px;" @click="artOk(artInfo.id)">{{id?'更新':'提交'}}</a-button>
            <a-button type="primary" @click="addCancel">取消</a-button>
        </a-form-model-item>
        </a-form-model>
    </a-card>
</div>
</template>
<script>

import URL from '@/plugin/axios'
import Editor from '../editor/index'
export default {
  components: { Editor },
  props: ['id'],
  data () {
    return {
      artInfo: {
        id: 0,
        cid: undefined,
        desc: '',
        title: '',
        content: '',
        img: ''
      },
      Catelist: [],
      upUrl: URL + '/upload',
      headers: {},
      artInfoRules: {
        title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
        cid: [{ required: true, message: '请输入分类', trigger: 'change' }],
        desc: [{ required: true, message: '请输入描述', trigger: 'blur' }, { max: 120, message: '描述最多120字', trigger: 'change' }],
        content: [{ required: true, message: '请输入文字内容', trigger: 'blur' }],
        img: [{ required: true, message: '请选择文章缩略图', trigger: 'blur' }]
      }

    }
  },
  created () {
    this.getCatelist()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
    if (this.id) {
      this.getArtInfo(this.id)
    }
  },
  methods: {
    // 获取文章信息
    async getArtInfo (id) {
      const { data: res } = await this.$http.get(`article/info/${id}`)
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      if (!res.data) {
        return this.$message.error('No data returned from API.')
      }
      this.artInfo = res.data
      this.artInfo.id = res.data.ID
      this.artInfo.content = res.data.content
    },
    // 获取分类列表
    async getCatelist () {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) {
        this.$message.error(res.message)
        return
      }
      if (!res.data) {
        this.$message.error('No data returned from API.')
        return
      }
      this.Catelist = res.data
    },
    Catechange (value) {
      this.artInfo.cid = value
    },
    // 上传图片
    upChange (info) {
      if (info.file.status !== 'uploading') {
        // console.log(info.file, info.fileList)
      }
      if (info.file.status === 'done') {
        this.$message.success(`${info.file.name} 上传成功`)
        const imgUrl = info.file.response.url
        this.artInfo.img = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error(`${info.file.name} 上传失败.`)
      }
    },
    // 提交&&更新文章
    artOk (id) {
      this.$refs.artInfoRef.validate(async (valid) => {
        if (id === 0) {
          const { data: res } = await this.$http.post('article/add', this.artInfo)
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$router.push('artlist')
          this.$message.success('文章添加成功')
        } else {
          const { data: res } = await this.$http.put(`article/${id}`, this.artInfo)
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$router.push('/admin/artlist')
          this.$message.success('文章更新成功')
        }
      })
    },
    addCancel () {
      this.$refs.artInfoRef.resetFields()
      if (this.id) {
        this.$router.push('/admin/artlist')
      }
    }
  }
}
</script>
