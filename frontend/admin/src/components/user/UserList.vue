<template>
  <div>
    <h3>用户列表页面</h3>
    <a-card class="userCard">
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search placeholder="输入用户名进行搜索" enter-button />
        </a-col>
        <a-col :span="4">
          <a-button type="primary">搜索</a-button>
        </a-col>
      </a-row>
      <a-table
        bordered
        rowKey="username"
        :columns="columns"
        :pagination="paginationOption"
        :dataSource="userlist"
      >
      <span slot="role" slot-scope="role">
        <a-tag v-if="role == '1'" color="red">
         <span style="font-size: 16px;">管理员</span>
        </a-tag>
        <a-tag v-if="role != '1'" color="green">
          <span style="font-size: 17px;">订阅者</span>
        </a-tag>
      </span>
      <template slot="action">
        <div class="actionSlot">
          <a-button type="primary" style="margin-right: 20px;">编辑</a-button>
          <a-button type="danger">删除</a-button>
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
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    scopedSlots: { customRender: 'role' },
    align: 'center'
  },
  {
    title: '操作',
    width: '20%',
    key: 'action',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  data () {
    return {
      paginationOption: {
        pageSizeOptions: ['5', '10', '20'],
        total: 0,
        defaultCurrent: 1,
        defaultPageSize: 5,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
        onChange: (current, pageSize) => {
          this.paginationOption.defaultCurrent = current
          this.paginationOption.defaultPageSize = pageSize
          this.getUserlist()
        },
        onShowSizeChange: (current, size) => {
          this.paginationOption.defaultCurrent = current
          this.paginationOption.defaultPageSize = size
          this.getUserlist()
        }
      },
      userlist: [],
      columns
    }
  },
  created () {
    this.getUserlist()
  },
  methods: {
    async getUserlist () {
      try {
        const { data: res } = await this.$http.get('users', {
          params: {
            pagesize: this.paginationOption.defaultPageSize,
            pagenum: this.paginationOption.defaultCurrent
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
        this.userlist = res.data
        this.paginationOption.total = res.total
        // console.log(res)
      } catch (error) {
        // console.error('Error fetching user list:', error)
        this.$message.error('Failed to fetch user list.')
      }
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
