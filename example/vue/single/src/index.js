import Button from '../packages/button/index.js'

const install = function (Vue, opts = {}) {
    Vue.component(Button.name,Button)
}

/* istanbul ignore if */
if (typeof window !== 'undefined' && window.Vue) {
    install(window.Vue);
}

export default { version:'1.0.0',install,Button }