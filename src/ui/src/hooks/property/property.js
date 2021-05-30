import { reactive, toRef, watch } from '@vue/composition-api'
import propertyService from '@/service/property/property'
export default function (options = {}) {
  const state = reactive({
    result: []
  })
  const refresh = async (value) => {
    if (!value.bk_obj_id) return
    state.result = await propertyService.find(value)
  }
  watch(() => options, refresh, { immediate: true, deep: true })
  return [toRef(state, 'result'), { refresh }]
}
