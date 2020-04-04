import Model from './Model.js'
import Vue from 'vue'

export default class ReactiveModel extends Model {
    constructor ({ request, primaryKey }) {
        super({ request, primaryKey })
        this.listeners = {}
        this.reactiveVM = new Vue({
            data: {
                state: {},
                changedState: {}
            }
        })
    }

    addStateListener (key, handler) {
        if (this.listeners.hasOwnProperty(key)) {
            this.listeners[key].push(handler)
        } else {
            this.listeners[key] = [handler]
        }
    }

    removeStateListener (key, handler) {
        if (!this.listeners.hasOwnProperty(key)) {
            return
        }
        const listeners = this.listeners[key]
        if (handler) {
            const index = listeners.indexOf(handler)
            if (index > -1) {
                listeners.splice(index, 1)
            }
        }
        listeners.splice(0)
    }

    flushCallbacks (key) {
        (this.listeners[key] || []).forEach(callback => callback())
    }

    defineReactive (state) {
        for (const key in state) {
            if (this.reactiveVM.state.hasOwnProperty(key)) {
                console.warn(`Can not redefine reactive key "${key}", already exist.`)
                continue
            }
            Vue.set(this.reactiveVM.state, key, state[key])
            Vue.set(this.reactiveVM.changedState, key, false)
            this.reactiveVM.$watch(() => this.reactiveVM.state[key], () => {
                this.flushCallbacks(key)
            })
            Object.defineProperty(this, key, {
                get () {
                    return this.reactiveVM.state[key]
                },
                set (val) {
                    this.reactiveVM.state[key] = val
                }
            })
        }
    }
}
