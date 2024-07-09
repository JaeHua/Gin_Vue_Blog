<template>
<div>
    <a-card>
        <h3>{{id?'编辑文章':'新增文章'}}</h3>
        <a-form-model v-model="artInfo">
        <a-form-model-item label="文章标题">
            <a-input style="width: 200px;" v-model="artInfo.title"></a-input>
        </a-form-model-item>
        <a-form-model-item label="文章分类">
            <a-select style="width: 200px;" @change="Catechange" v-model="artInfo.cid" placeholder="请选择分类">
                <a-select-option v-for="item in Catelist" :key="item.id" :value="item.id">{{item.name}}</a-select-option>
            </a-select>
        </a-form-model-item>
        <a-form-model-item label="文章描述">
            <a-input type="textarea" v-model="artInfo.desc"></a-input>
        </a-form-model-item>
        <a-form-model-item label="文章缩略图">
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
        </a-form-model-item>
        <a-form-model-item label="文章内容">
            <a-input type="textarea" v-model="artInfo.content"></a-input>
        </a-form-model-item>
        <a-form-model-item>
            <a-button type="danger" style="margin-right: 15px;">提交</a-button>
            <a-button type="primary">取消</a-button>
        </a-form-model-item>
        </a-form-model>
    </a-card>
</div>
</template>
<script>
import URL from '@/plugin/axios'
export default {
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
      headers: {}
    }
  },
  created () {
    this.getCatelist()
    this.headers = { Authorization: `Bearer ${window.sessionStorage.getItem('token')}` }
    console.log('id' + this.id)
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
        const imgUrl = info.file.response.Url
        this.artInfo.img = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error(`${info.file.name} 上传失败.`)
      }
    }
  }
}
</script>
