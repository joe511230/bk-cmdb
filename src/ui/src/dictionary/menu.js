import {
    MENU_INDEX,
    MENU_BUSINESS,
    MENU_RESOURCE,
    MENU_MODEL,
    MENU_ANALYSIS,
    MENU_BUSINESS_HOST,
    MENU_BUSINESS_SERVICE,
    MENU_BUSINESS_ADVANCED,
    MENU_BUSINESS_CLOUD_DISCOVER,
    MENU_RESOURCE_EVENTPUSH,
    MENU_RESOURCE_MANAGEMENT,
    MENU_MODEL_MANAGEMENT,
    MENU_MODEL_TOPOLOGY,
    MENU_MODEL_TOPOLOGY_NEW,
    MENU_MODEL_BUSINESS_TOPOLOGY,
    MENU_MODEL_ASSOCIATION,
    MENU_ANALYSIS_AUDIT,
    MENU_ANALYSIS_OPERATION
} from './menu-symbol'
import {
    businessViews,
    resourceViews,
    modelViews,
    analysisViews
} from '@/views'

const getSubmenu = (views, symbol, pathPrefix = '') => {
    const submenuViews = views.filter(view => {
        return view.meta.menu.parent === symbol
    })
    const submenu = submenuViews.map(view => {
        const menu = view.meta.menu
        return {
            id: Symbol(menu.i18n),
            i18n: menu.i18n,
            route: getMenuRoute(view, symbol, pathPrefix)
        }
    })
    return submenu
}

const getMenuRoute = (views, symbol, pathPrefix = '') => {
    const menuView = Array.isArray(views)
        ? views.find(view => view.name === symbol)
        : views
    if (menuView) {
        return {
            name: menuView.name,
            path: `/${pathPrefix}/${menuView.path}`,
            available: menuView.meta.available
        }
    }
    return {}
}

const menus = [{
    id: MENU_INDEX,
    i18n: '首页'
}, {
    id: MENU_BUSINESS,
    i18n: '业务',
    menu: [{
        id: MENU_BUSINESS_HOST,
        i18n: '主机',
        icon: 'icon-cc-host',
        submenu: getSubmenu(businessViews, MENU_BUSINESS_HOST, 'business')
    }, {
        id: MENU_BUSINESS_SERVICE,
        i18n: '服务',
        icon: 'icon-cc-template-management',
        submenu: getSubmenu(businessViews, MENU_BUSINESS_SERVICE, 'business')
    }, {
        id: MENU_BUSINESS_ADVANCED,
        i18n: '高级功能',
        icon: 'icon-cc-nav-advanced-features',
        submenu: getSubmenu(businessViews, MENU_BUSINESS_ADVANCED, 'business')
    }, {
        id: MENU_BUSINESS_CLOUD_DISCOVER,
        i18n: '云资源发现',
        icon: 'icon-cc-nav-advanced-features',
        route: getMenuRoute(businessViews, MENU_BUSINESS_CLOUD_DISCOVER, 'business')
    }]
}, {
    id: MENU_RESOURCE,
    i18n: '资源',
    menu: [{
        id: MENU_RESOURCE_MANAGEMENT,
        i18n: '资源目录',
        icon: 'icon-cc-square',
        route: getMenuRoute(resourceViews, MENU_RESOURCE_MANAGEMENT, 'resource')
    }, {
        id: MENU_RESOURCE_EVENTPUSH,
        i18n: '事件订阅',
        icon: 'icon-cc-nav-subscription',
        route: getMenuRoute(resourceViews, MENU_RESOURCE_EVENTPUSH, 'resource')
    }]
}, {
    id: MENU_MODEL,
    i18n: '模型',
    menu: [{
        id: MENU_MODEL_MANAGEMENT,
        i18n: '模型管理',
        icon: 'icon-cc-nav-model-02',
        route: getMenuRoute(modelViews, MENU_MODEL_MANAGEMENT, 'model')
    }, {
        id: MENU_MODEL_TOPOLOGY,
        i18n: '模型拓扑',
        icon: 'icon-cc-nav-model-topo',
        route: getMenuRoute(modelViews, MENU_MODEL_TOPOLOGY, 'model')
    }, {
        id: MENU_MODEL_TOPOLOGY_NEW,
        i18n: '模型关系',
        icon: 'icon-cc-nav-model-topo',
        route: getMenuRoute(modelViews, MENU_MODEL_TOPOLOGY_NEW, 'model')
    }, {
        id: MENU_MODEL_BUSINESS_TOPOLOGY,
        i18n: '业务层级',
        icon: 'icon-cc-tree',
        route: getMenuRoute(modelViews, MENU_MODEL_BUSINESS_TOPOLOGY, 'model')
    }, {
        id: MENU_MODEL_ASSOCIATION,
        i18n: '关联类型',
        icon: 'icon-cc-nav-associated',
        route: getMenuRoute(modelViews, MENU_MODEL_ASSOCIATION, 'model')
    }]
}, {
    id: MENU_ANALYSIS,
    i18n: '运营分析',
    menu: [{
        id: MENU_ANALYSIS_AUDIT,
        i18n: '操作审计',
        icon: 'icon-cc-nav-audit-02',
        route: getMenuRoute(analysisViews, MENU_ANALYSIS_AUDIT, 'analysis')
    }, {
        id: MENU_ANALYSIS_OPERATION,
        i18n: '运营统计',
        icon: 'icon-cc-nav-associated',
        route: getMenuRoute(analysisViews, MENU_ANALYSIS_OPERATION, 'analysis')
    }]
}]

// 移除未被激活的menu
;(() => {
    menus.forEach(top => {
        if (top.hasOwnProperty('menu')) {
            top.menu.forEach(menu => {
                if (menu.hasOwnProperty('submenu')) {
                    menu.submenu = menu.submenu.filter(submenu => submenu.route.available)
                }
            })
            top.menu = top.menu.filter(menu => {
                if (menu.hasOwnProperty('route')) {
                    return menu.route.available
                }
                return menu.submenu.length
            })
        }
    })
})()

export default menus
