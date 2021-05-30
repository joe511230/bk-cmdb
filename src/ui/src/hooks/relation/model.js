import { ref, watch } from '@vue/composition-api'
import modelRelationService from '@/service/relation/model'
export default function (modelId) {
  const relations = ref([])
  const refresh = async (value) => {
    if (!value) return
    relations.value = await modelRelationService.findAll(value)
  }
  watch(modelId, refresh, { immediate: true })
  return [relations, refresh]
}
