<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="活动id:" prop="activityId">
          <el-input v-model.number="formData.activityId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="填写的内推码:" prop="code">
          <el-input v-model="formData.code" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="参与状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户id:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:" prop="note">
          <el-input v-model="formData.note" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="评分:" prop="score">
          <el-input v-model.number="formData.score" :clearable="true" placeholder="请输入" />
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
  name: 'ActivityRecord'
}
</script>

<script setup>
import {
  createActivityRecord,
  updateActivityRecord,
  findActivityRecord
} from '@/api/activityRecord'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            activityId: 0,
            code: '',
            status: '',
            userId: 0,
            note: '',
            score: 0,
        })
// 验证规则
const rule = reactive({
               activityId : [{
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
      const res = await findActivityRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reactivityRecord
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
               res = await createActivityRecord(formData.value)
               break
             case 'update':
               res = await updateActivityRecord(formData.value)
               break
             default:
               res = await createActivityRecord(formData.value)
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
