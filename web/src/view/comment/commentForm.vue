<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="内容:" prop="content">
          <el-input v-model="formData.content" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="描述(a评论了xxx):" prop="describe">
          <el-input v-model="formData.describe" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否匿名（0否，1是）:" prop="isAnonymous">
          <el-switch v-model="formData.isAnonymous" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="评分:" prop="rate">
          <el-input v-model.number="formData.rate" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="引用回复:" prop="referId">
          <el-input v-model.number="formData.referId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="标签(已毕业、社团成员):" prop="tags">
          <el-input v-model="formData.tags" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="评论对象id:" prop="toId">
          <el-input v-model.number="formData.toId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="评论对象类别:" prop="toType">
          <el-input v-model.number="formData.toType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发表者id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="voteDownNum字段:" prop="voteDownNum">
          <el-input v-model.number="formData.voteDownNum" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="voteUpNum字段:" prop="voteUpNum">
          <el-input v-model.number="formData.voteUpNum" :clearable="true" placeholder="请输入" />
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
  name: 'Comment'
}
</script>

<script setup>
import {
  createComment,
  updateComment,
  findComment
} from '@/api/comment'

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
            describe: '',
            isAnonymous: false,
            rate: 0,
            referId: 0,
            tags: '',
            toId: 0,
            toType: 0,
            userId: 0,
            voteDownNum: 0,
            voteUpNum: 0,
        })
// 验证规则
const rule = reactive({
               content : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               toId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               toType : [{
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
      const res = await findComment({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recomment
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
               res = await createComment(formData.value)
               break
             case 'update':
               res = await updateComment(formData.value)
               break
             default:
               res = await createComment(formData.value)
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
