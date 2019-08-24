import HelloWorld from './HelloWorld.vue'

HelloWorld.install = function(Vue) {
    Vue.component(HelloWorld.name, HelloWorld);
};

export default HelloWorld;
