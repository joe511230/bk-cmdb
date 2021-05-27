import { ref } from '@vue/composition-api'
import { importHost } from '@/service/host/import-export'
const file = ref(null)
const state = ref(null)

const submit = async (params) => {
  debugger
  if (!file.value) return
  state.value = 'pending'
  try {
    const form = new FormData()
    form.append('file', file.value)
    form.append('params', JSON.stringify(params))
    await importHost(form)
    state.value = 'success'
    return true
  } catch (error) {
    console.error(error)
    state.value = 'error'
  }
}

const change = (event) => {
  const { files: [userFile] } = event.target
  file.value = userFile
}

const clear = () => {
  file.value = null
  state.value = null
}
export default function () {
  return [{ file, state }, { change, submit, clear }]
}
