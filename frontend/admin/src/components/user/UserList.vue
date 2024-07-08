<template>
  <div>
    <h3>用户列表页面</h3>
    <a-card class="userCard">
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search allowClear v-model="queryParam.username" placeholder="输入用户名进行搜索" enter-button  @change="getUserlist" />
        </a-col>
        <a-col :span="4">
          <a-button  type="primary" @click="addUservisible=true">新增</a-button>
        </a-col>
      </a-row>
      <a-table
        bordered
        rowKey="username"
        :columns="columns"
        :pagination="pagination"
        :dataSource="userlist"
        @change="handleTableChange"
      >
      <span slot="role" slot-scope="role">
        <a-tag v-if="role == '1'" color="red">
         <span style="font-size: 16px;">管理员</span>
        </a-tag>
        <a-tag v-if="role != '1'" color="green">
          <span style="font-size: 17px;">订阅者</span>
        </a-tag>
      </span>
      <template slot="action" slot-scope="data">
        <div class="actionSlot">
          <a-button type="primary" icon="edit" @click="editUser(data.ID)" style="margin-right: 20px;">编辑</a-button>
          <a-button type="danger" icon="delete" @click="deleteUser(data.ID)" style="margin-right: 20px;">删除</a-button>
          <a-button type="info" icon="scissor" @click="resetUser(data.ID)">重置密码</a-button>
        </div>
      </template>
      </a-table>
    </a-card>
    <!-- 新增用户 -->
    <a-modal
    title="新增用户"
    :visible="addUservisible"
    @ok="addUserOk"
    @cancel="addUserCancel"
    closable
    :destroyOnClose="true"
  >
    <a-form-model :model="newUserInfo" :rules="adduserRules" ref="userRef">
      <a-form-model-item label="用户名" prop="username">
        <a-input v-model="newUserInfo.username"></a-input>
      </a-form-model-item>
      <a-form-model-item hasFeedback label="密码" prop="password">
        <a-input-password  v-model="newUserInfo.password"></a-input-password>
      </a-form-model-item>
      <a-form-model-item hasFeedback label="确认密码" prop="checkpass">
        <a-input-password  v-model="newUserInfo.checkpass"></a-input-password>
      </a-form-model-item>
    </a-form-model>
    </a-modal>
    <!-- 编辑用户 -->
    <a-modal
    title="编辑用户"
    :visible="editUservisible"
    @ok="editUserOk"
    @cancel="editUserCancel"
    closable
    :destroyOnClose="true"
  >
    <a-form-model :model="userInfo" :rules="userRules" ref="userRef">
      <a-form-model-item label="用户名" prop="username">
        <a-input v-model="userInfo.username"></a-input>
      </a-form-model-item>
      <a-form-model-item label="是否为管理员"  prop="role">
        <a-switch checked-children="是" un-checked-children="否"  :defaultChecked="Boolean((this.userInfo.role===1))"  @change="adminChange" />
      </a-form-model-item>
    </a-form-model>
    </a-modal>
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
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        total: 0,
        pageSize: 5,
        current: 1,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`
      },
      userlist: [],
      columns,
      queryParam: {
        username: '',
        pagesize: 5,
        pagenum: 1
      },
      visible: false,
      addUservisible: false,
      editUservisible: false,
      userInfo: {
        id: 0,
        username: '',
        password: '',
        checkpass: '',
        role: 2
      },
      newUserInfo: {
        id: 0,
        username: '',
        password: '',
        checkpass: '',
        role: 2
      },
      userRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名必须在4到12个字符之间', trigger: 'blur' }
        ],
        password: [{
          validator: (rule, value, callback) => {
            if (this.userInfo.password === '') { callback(new Error('请输入密码')) }
            if ([...this.userInfo.password].length < 6 || [...this.userInfo.password].length > 20) { callback(new Error('密码应当在6到20位之间')) } else {
              callback()
            }
          },
          trigger: 'blur'
        }],
        checkpass: [{
          validator: (rule, value, callback) => {
            if (this.userInfo.checkpass === '') { callback(new Error('请再次输入密码')) }
            if (this.userInfo.checkpass !== this.userInfo.password) { callback(new Error('密码不一致')) } else {
              callback()
            }
          },
          trigger: 'blur'
        }]
      },
      adduserRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名必须在4到12个字符之间', trigger: 'blur' }
        ],
        password: [{
          validator: (rule, value, callback) => {
            if (this.newUserInfo.password === '') { callback(new Error('请输入密码')) }
            if ([...this.newUserInfo.password].length < 6 || [...this.newUserInfo.password].length > 20) { callback(new Error('密码应当在6到20位之间')) } else {
              callback()
            }
          },
          trigger: 'blur'
        }],
        checkpass: [{
          validator: (rule, value, callback) => {
            if (this.newUserInfo.checkpass === '') { callback(new Error('请再次输入密码')) }
            if (this.newUserInfo.checkpass !== this.newUserInfo.password) { callback(new Error('密码不一致')) } else {
              callback()
            }
          },
          trigger: 'blur'
        }]
      }
    }
  },
  created () {
    this.getUserlist()
  },
  methods: {
    async getUserlist () {
      const { data: res } = await this.$http.get('users', {
        params: {
          username: this.queryParam.username,
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
      this.userlist = res.data
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
      this.getUserlist()
    },
    // 删除用户
    deleteUser (id) {
      this.$confirm({
        title: '提示:确定删除该用户吗?',
        content: '一旦删除，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`user/${id}`)
          console.log(res)
          if (res.status !== 200) {
            this.$message.error(res.message)
          }
          this.$message.success('删除成功')
          this.getUserlist()
        },
        onCancel: async () => {
          this.$message.info('已取消删除')
        }
      })
    },
    resetUser (id) {
      this.$confirm({
        title: '提示:确定重置密码?',
        content: '一旦重置，无法恢复',
        onOk: async () => {
          const { data: res } = await this.$http.put(`userinfo/${id}`)
          console.log(res)
          if (res.status !== 200) {
            this.$message.error(res.message)
          }
          this.$message.success('重置密码成功')
          this.getUserlist()
        },
        onCancel: async () => {
          this.$message.info('已取消重置密码')
        }
      })
    },
    // 新增用户
    addUserOk () {
      this.$refs.userRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')

        const { data: res } = await this.$http.post('user/add', { username: this.newUserInfo.username, password: this.newUserInfo.password, role: this.newUserInfo.role })
        this.$message.success('用户添加成功')
        if (res.status !== 200) {
          return this.$message.error(res.message)
        }
      })
      this.addUservisible = false
      this.getUserlist()
    },
    addUserCancel () {
      this.$refs.userRef.resetFields()
      this.addUservisible = false
      this.$message.info('编辑已取消')
    },
    adminChange (checked) {
      if (checked) {
        this.userInfo.role = 1
      } else {
        this.userInfo.role = 2
      }
    },
    async editUser (id) {
      this.editUservisible = true
      const { data: res } = await this.$http.get(`user/${id}`)
      this.userInfo = res.data
      this.userInfo.id = id
    },
    editUserCancel () {
      this.$refs.userRef.resetFields()
      this.editUservisible = false
      this.$message.info('编辑已取消')
    },
    editUserOk () {
      this.$refs.userRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, { username: this.userInfo.username, role: this.userInfo.role })
        if (res.status !== 200) {
          return this.$message.error(res.message)
        }
        this.$message.success('用户更新成功')
        this.editUservisible = false
        this.getUserlist()
      })
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
