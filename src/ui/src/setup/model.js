import Vue from 'vue'

import Directory from '@/models/Directory'

const defineMap = {
    [Directory.name]: Directory
}
const models = {}

Vue.prototype.CMDB_MODEL = {
    get (name) {
        if (models.hasOwnProperty(name)) {
            return models[name]
        }
        if (defineMap.hasOwnProperty(name)) {
            const instance = models[name] = new defineMap[name]()
            return instance
        }
        return null
    }
}
