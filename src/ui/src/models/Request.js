import API from '@/api'

function withParams (method) {
    return ['post', 'put', 'patch'].includes(method)
}

function convertArguments (args) {
    return [].slice.apply(args)
}

const key = Symbol('options')

export default class Request {
    constructor (options) {
        this[key] = options
    }

    getOptions () {
        const args = convertArguments(arguments)
        const type = args[0]
        return this[key][type](args.slice(1))
    }

    request (options) {
        const { url, method, params, config, resolver } = options
        let promise
        if (withParams(method)) {
            promise = API[method](url, params, config)
        } else {
            promise = API[method](url, config)
        }
        if (typeof resolver === 'function') {
            return promise.then(response => resolver({
                input: params,
                output: response
            }))
        }
        return promise
    }

    async read (params) {
        try {
            const options = this.getOptions('read', params)
            return await this.request({ ...options, params })
        } catch (e) {
            return Promise.reject(e)
        }
    }

    async create (params) {
        try {
            const options = this.getOptions('create', params)
            return await this.request({ ...options, params })
        } catch (e) {
            return Promise.reject(e)
        }
    }

    async update (id, params) {
        try {
            const options = this.getOptions('update', id, params)
            return await this.request({ ...options, params })
        } catch (e) {
            return Promise.reject(e)
        }
    }

    async delete (id) {
        try {
            const options = this.getOptions('delete', id)
            return await this.request(options)
        } catch (e) {
            return Promise.reject(e)
        }
    }
}
