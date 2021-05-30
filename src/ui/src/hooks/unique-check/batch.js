import { ref, watch } from '@vue/composition-api'
import uniqueCheckService from '@/service/unique-check'
export default function (models) {
  const uniqueChecks = ref([])
  const refresh = async (value) => {
    if (!value.length) return
    uniqueChecks.value = await uniqueCheckService.findMany(value)
  }
  watch(models, refresh, { immediate: true })
  return [uniqueChecks, { refresh }]
}
