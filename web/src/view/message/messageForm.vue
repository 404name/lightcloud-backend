<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="动态内容:" prop="content">
          <el-input v-model="formData.content" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="封面:" prop="cover">
          <el-input v-model="formData.cover" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="消息发起id:" prop="fromId">
          <el-input v-model.number="formData.fromId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="消息发起方:" prop="fromType">
          <el-input v-model.number="formData.fromType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否置顶（0否, 1是）:" prop="isTop">
          <el-switch v-model="formData.isTop" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="标题(AXX了你/通知):" prop="title">
          <el-input v-model="formData.title" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类别(通知，动态，):" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Message'
}
</script>

<script setup>
import {
  createMessage,
  updateMessage,
  findMessage
} from '@/api/message'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            content: '',
            cover: '',
            fromId: 0,
            fromType: 0,
            isTop: false,
            title: '',
            type: '',
            userId: 0,
        })
// 验证规则
const rule = reactive({
               content : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cover : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               fromId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               fromType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMessage({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remessage
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createMessage(formData.value)
               break
             case 'update':
               res = await updateMessage(formData.value)
               break
             default:
               res = await createMessage(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
