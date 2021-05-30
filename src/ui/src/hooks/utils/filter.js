import { toRef, reactive, watch, computed } from '@vue/composition-api'
import debounce from 'lodash.debounce'

export default function ({ list, keyword, target, available }) {
  const state = reactive({
    result: []
  })
  const availableList = computed(() => list.value.filter(available))
  const handler = (value) => {
    if (!value) {
      state.result = availableList.value
      return
    }
    const regexp = new RegExp(value, 'ig')
    state.result = availableList.value.filter(item => regexp.test(item[target]))
  }
  const filter = debounce(handler, 300, { leading: false, trailing: true })
  watch(keyword, filter)
  watch(list, () => handler(), { immediate: true })
  return [toRef(state, 'result'), { filter }]
}
