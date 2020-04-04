import Vue from 'vue'
import Request from './Request.js'
export default class Model extends Request {
    constructor ({ request, primaryKey }) {
        super(request)

        this.primaryKey = primaryKey
        this.data = null
    }

    getAll () {
        return this.data
    }

    get (id) {
        return this.data.find(target => target[this.primaryKey] === id)
    }

    getIndexOf (index) {
        if (index >= this.data.length) {
            return null
        }
        return this.data[index]
    }

    async read (params) {
        try {
            if (this.data) {
                return Promise.resolve(this.data)
            }
            const data = await super.read(params)
            this.updateData({
                newValue: data,
                type: 'read'
            })
            return Promise.resolve(data)
        } catch (e) {
            Promise.reject(e)
        }
    }

    async create (params) {
        try {
            const datum = await super.create(params)
            this.data.push(datum)
            return Promise.resolve(datum)
        } catch (e) {
            return Promise.reject(e)
        }
    }

    async update (id, params) {
        try {
            const newValue = await super.update(id, params)
            this.updateData({
                id,
                newValue,
                type: 'update'
            })
            return Promise.resolve(newValue)
        } catch (e) {
            return Promise.reject(e)
        }
    }

    async delete (id) {
        try {
            const result = await super.delete(id)
            this.updateData({
                id,
                type: 'delete'
            })
            return Promise.resolve(result)
        } catch (e) {
            return Promise.reject(e)
        }
    }

    updateData ({ type, id, newValue }) {
        if (type === 'read') {
            Vue.set(this, 'data', newValue)
            return
        }

        const index = this.data.findIndex(old => old[this.primaryKey] === id)
        if (index === -1) {
            return
        }
        if (type === 'update') {
            this.data.splice(index, 1, newValue)
        } else if (type === 'delete') {
            this.data.splice(index, 1)
        }
    }
}
