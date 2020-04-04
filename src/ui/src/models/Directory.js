import ReactiveModel from './ReactiveModel.js'

export default class Directory extends ReactiveModel {
    constructor () {
        super({
            primaryKey: 'bk_module_id',
            request: {
                read: () => ({
                    url: 'findmany/resource/directory',
                    method: 'post',
                    config: {
                        requestId: Symbol('directory.read')
                    },
                    resolver: data => {
                        return data.output.info
                    }
                }),
                create: () => ({
                    url: 'create/resource/directory',
                    method: 'post',
                    config: {
                        requestId: Symbol('directory.create')
                    }
                }),
                update: id => ({
                    url: `update/resource/directory/${id}`,
                    method: 'put',
                    config: {
                        requestId: Symbol('directory.update')
                    }
                }),
                delete: id => ({
                    url: `delete/resource/directory/${id}`,
                    method: 'delete',
                    config: {
                        requestId: Symbol('directory.delete')
                    }
                })
            }
        })
    }
}
