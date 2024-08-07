import Vue from 'vue'
import {
  Button,
  FormModel,
  Input,
  Icon,
  message,
  Layout,
  Menu,
  Card,
  Table,
  Row,
  Col,
  Tag,
  ConfigProvider,
  Modal,
  Select,
  Switch,
  Upload
} from 'ant-design-vue'
message.config({
  top: '100px',
  duration: 2,
  maxCount: 3
})
Vue.prototype.$message = message
Vue.prototype.$confirm = Modal.confirm

Vue.use(Button)
Vue.use(FormModel)
Vue.use(Input)
Vue.use(Icon)
Vue.use(Layout)
Vue.use(Menu)
Vue.use(Table)
Vue.use(Card)
Vue.use(Row)
Vue.use(Col)
Vue.use(Tag)
Vue.use(ConfigProvider)
Vue.use(Modal)
Vue.use(Select)
Vue.use(Switch)
Vue.use(Upload)
