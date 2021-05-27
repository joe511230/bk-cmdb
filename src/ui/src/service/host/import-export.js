import http from '@/api'
// eslint-disable-next-line arrow-body-style
export const importHost = (form, config) => {
  return http.post(`${window.API_HOST}hosts/update`, form, config)
}

// eslint-disable-next-line arrow-body-style
export const exportHost = (params, config) => {
  return http.post(`${window.API_HOST}hosts/export`, params, config)
}

export default {
  import: importHost,
  export: exportHost
}
