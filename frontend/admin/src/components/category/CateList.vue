<template>
    <div>
      <h3>分类列表页面</h3>
      <a-card class="CateCard">
        <a-row :gutter="20">
          <a-col :span="4">
            <a-button  type="primary" @click="addCatevisible=true">新增</a-button>
          </a-col>
        </a-row>
        <a-table
          bordered
          rowKey="name"
          :columns="columns"
          :pagination="pagination"
          :dataSource="Catelist"
          @change="handleTableChange"
        >
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" @click="editCate(data.id)" style="margin-right: 20px;">编辑</a-button>
            <a-button type="danger" icon="delete" @click="deleteCate(data.id)" style="margin-right: 20px;">删除</a-button>
          </div>
        </template>
        </a-table>
      </a-card>
      <!-- 新增分类 -->
      <a-modal
      title="新增分类"
      :visible="addCatevisible"
      @ok="addCateOk"
      @cancel="addCateCancel"
      closable
      :destroyOnClose="true"
    >
      <a-form-model :model="newCateInfo" :rules="addCateRules" ref="CateRef">
        <a-form-model-item label="分类名" prop="name">
          <a-input v-model="newCateInfo.name"></a-input>
        </a-form-model-item>
      </a-form-model>
      </a-modal>
      <!-- 编辑分类 -->
      <a-modal
      title="编辑分类"
      :visible="editCatevisible"
      @ok="editCateOk"
      @cancel="editCateCancel"
      closable
      :destroyOnClose="true"
    >
      <a-form-model :model="CateInfo" :rules="CateRules" ref="CateRef">
        <a-form-model-item label="分类名" prop="name">
          <a-input v-model="CateInfo.name"></a-input>
        </a-form-model-item>
      </a-form-model>
      </a-modal>
    </div>
  </template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    width: '5%',
    key: 'id',
    align: 'center'
  },
  {
    title: '分类名',
    dataIndex: 'name',
    width: '20%',
    key: 'name',
    align: 'center'
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
      Catelist: [],
      columns,
      queryParam: {
        Catename: '',
        pagesize: 5,
        pagenum: 1
      },
      visible: false,
      addCatevisible: false,
      editCatevisible: false,
      CateInfo: {
        id: 0,
        name: ''
      },
      newCateInfo: {
        name: ''
      },
      CateRules: {
        name: [
          { required: true, message: '请输入分类名', trigger: 'blur' }
        ]
      },
      addCateRules: {
        name: [
          { required: true, message: '请输入分类名', trigger: 'blur' }
        ]
      }
    }
  },
  created () {
    this.getCatelist()
  },
  methods: {
    // 获取分类列表
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
      this.getCatelist()
    },
    // 删除分类
    deleteCate (id) {
      this.$confirm({
        title: '提示:确定删除该分类吗?',
        content: '一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`category/${id}`)
          console.log(res)
          if (res.status !== 200) {
            this.$message.error(res.message)
          }
          this.$message.success('删除成功')
          this.getCatelist()
        },
        onCancel: async () => {
          this.$message.info('已取消删除')
        }
      })
    },
    // 新增分类
    addCateOk () {
      this.$refs.CateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')

        const { data: res } = await this.$http.post('category/add', { name: this.newCateInfo.name })
        this.$message.success('分类添加成功')
        if (res.status !== 200) {
          return this.$message.error(res.message)
        }
        this.addCatevisible = false
        this.getCatelist()
      })
      this.addCatevisible = false
      this.getCatelist()
    },
    addCateCancel () {
      this.$refs.CateRef.resetFields()
      this.addCatevisible = false
      this.$message.info('编辑已取消')
    },
    // 编辑分类
    async editCate (id) {
      this.editCatevisible = true
      const { data: res } = await this.$http.get(`category/${id}`)
      this.CateInfo = res.data
      this.CateInfo.id = id
    },
    editCateCancel () {
      this.$refs.CateRef.resetFields()
      this.editCatevisible = false
      this.$message.info('编辑已取消')
    },
    editCateOk () {
      this.$refs.CateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`category/${this.CateInfo.id}`, { name: this.CateInfo.name })
        if (res.status !== 200) {
          return this.$message.error(res.message)
        }
        this.$message.success('分类更新成功')
        this.editCatevisible = false
        this.getCatelist()
      })
      this.editCatevisible = false
      this.getCatelist()
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
