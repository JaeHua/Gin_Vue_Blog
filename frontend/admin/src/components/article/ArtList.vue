<template>
    <div>
      <h3>文章列表页面</h3>
      <a-card class="ArtCard">
        <a-row :gutter="20">
          <a-col :span="6">
            <a-input-search allowClear v-model="queryParam.title" placeholder="输入文章标题进行搜索" enter-button  @change="getArtlist" />
          </a-col>
          <a-col :span="4">
            <a-button  type="primary" @click="$router.push('addart')">新增</a-button>
          </a-col>
          <a-col :span="6" :offset="1">
            <a-select  firstActiveValue defaultValue="选择分类" @change="Catechange" style="width: 120px;">
              <a-select-option :value="''">全部分类</a-select-option>
              <a-select-option  v-for="item in Catelist" :key="item.id" :value="item.id">{{item.name}}</a-select-option>
            </a-select>
          </a-col>
        </a-row>

        <a-table
          bordered
          rowKey="ID"
          :columns="columns"
          :pagination="pagination"
          :dataSource="Artlist"
          @change="handleTableChange"
        >

          <span slot="img" slot-scope="img" class="Artimg">
            <img :src="img" alt="Image"  />
          </span>

        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" @click="$router.push(`addart/${data.ID}`)" style="margin-right: 20px;">编辑</a-button>
            <a-button type="danger" icon="delete" @click="deleteArt(data.ID)" style="margin-right: 20px;">删除</a-button>
          </div>
        </template>
        </a-table>
      </a-card>
    </div>
  </template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '5%',
    key: 'id',
    align: 'center'
  },
  {
    title: '分类',
    dataIndex: 'Category.name',
    width: '10%',
    key: 'name',
    align: 'center'
  },
  {
    title: '文章标题',
    dataIndex: 'title',
    width: '15%',
    key: 'title',
    align: 'center'
  },
  {
    title: '文章描述',
    dataIndex: 'desc',
    width: '20%',
    key: 'desc',
    align: 'center'
  },
  {
    title: '缩略图',
    dataIndex: 'img',
    width: '20%',
    key: 'img',
    align: 'center',
    scopedSlots: { customRender: 'img' }

  },
  {
    title: '操作',
    width: '15%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  data () {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        total: 0,
        pageSize: 5,
        current: 1,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`
      },
      Artlist: [],
      Catelist: [],
      columns,
      queryParam: {
        title: '',
        pagesize: 5,
        pagenum: 1
      }
    }
  },
  created () {
    this.getArtlist()
    this.getCatelist()
  },
  methods: {
    // 获取文章列表
    async getArtlist () {
      const { data: res } = await this.$http.get('article', {
        params: {
          title: this.queryParam.title,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        this.$message.error(res.message)
        return
      }
      if (!res.data) {
        this.$message.error('No data returned from API.')
        return
      }
      this.Artlist = res.data
      this.pagination.total = res.total
      // console.log(res)
    },
    async getCatelist () {
      const { data: res } = await this.$http.get('category', {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        this.$message.error(res.message)
        return
      }
      if (!res.data) {
        this.$message.error('No data returned from API.')
        return
      }
      this.Catelist = res.data
      this.pagination.total = res.total
      // console.log(res)
    },
    // 更改分页
    handleTableChange (pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getArtlist()
    },
    // 删除用户
    deleteArt (id) {
      this.$confirm({
        title: '提示:确定删除该文章吗?',
        content: '一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`article/${id}`)
          console.log(res)
          if (res.status !== 200) {
            this.$message.error(res.message)
          }
          this.$message.success('删除成功')
          this.getArtlist()
        },
        onCancel: async () => {
          this.$message.info('已取消删除')
        }
      })
    },
    // 查询分类下的文章
    Catechange (value) {
      this.getCateArt(value)
    },
    async getCateArt (id) {
      const url = id ? `category/article/${id}` : 'article' // Adjust endpoint as needed
      const { data: res } = await this.$http.get(url, {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        this.$message.error(res.message)
        return
      }
      if (!res.data) {
        this.$message.error('No data returned from API.')
        return
      }
      this.Artlist = Array.isArray(res.data) ? res.data : []
      this.pagination.total = res.total || 0
    }
  }

}
</script>
  <style scoped>
  .actionSlot{
    display: flex;
    justify-content: center;
  }

  </style>
