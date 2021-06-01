import { computed } from '@vue/composition-api'

export default function (pendings = []) {
  const pending = computed(() => pendings.some(pending => pending.value))
  return pending
}
