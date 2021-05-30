import { reactive, toRefs } from '@vue/composition-api'
import useClone from '@/hooks/utils/clone'
const defaultState = {
  visible: false,
  title: '',
  step: 1,
  bk_obj_id: null,
  bk_biz_id: null,
  fields: [],
  status: null,
  available: () => true,
  submit: () => {},
  selection: [],
  relations: {}
}

const state = reactive(useClone(defaultState))

const setState = (newState) => {
  Object.assign(state, newState)
}
const resetState = () => setState(useClone(defaultState))

export default function () {
  return [toRefs(state), { setState, resetState }]
}
