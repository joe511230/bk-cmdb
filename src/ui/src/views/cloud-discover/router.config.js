import Meta from '@/router/meta'
import { MENU_BUSINESS, MENU_BUSINESS_CLOUD_DISCOVER } from '@/dictionary/menu-symbol'
import {
    C_CLOUD_DISCOVER,
    U_CLOUD_DISCOVER,
    D_CLOUD_DISCOVER,
    R_CLOUD_DISCOVER
} from '@/dictionary/auth'

export const OPERATION = {
    R_CLOUD_DISCOVER,
    C_CLOUD_DISCOVER,
    U_CLOUD_DISCOVER,
    D_CLOUD_DISCOVER
}

const path = 'cloud-discover'

export default {
    name: MENU_BUSINESS_CLOUD_DISCOVER,
    path: path,
    component: () => import('./index.vue'),
    meta: new Meta({
        menu: {
            owner: MENU_BUSINESS,
            i18n: '云资源发现'
        },
        auth: {
            operation: {},
            authScope: 'global'
        }
    })
}
