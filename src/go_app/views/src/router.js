import Vue from 'vue'
import Login from './components/Login.vue'
import SignUp from './components/SignUp.vue'
import Home from './components/Home.vue'
import WhatDone from './components/WhatDone.vue'
import Platform from './components/Platform.vue'
import Router from 'vue-router'
import Todo from './components/Todo.vue';
import TodoList from './components/TodoList.vue'
import Error401 from './components/errors/Error401.vue';
import Error500 from './components/errors/Error500.vue';

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login
    },
    {
      path: '/signUp',
      name: 'SignUp',
      component: SignUp
    },
    {
      path: '/platform',
      name: 'Platform',
      component: Platform,
    },
    {
      path: '/home',
      name: 'Home',
      component: Home
    },
    {
      path: '/whatDone/:action',
      name: 'WhatDone',
      component: WhatDone
    },
    {
      path: '/todoList',
      name: 'TodoList',
      component: TodoList
    },
    {
      path: '/todo/:action',
      name: 'Todo',
      component: Todo
    },
    {
      path: '/err401',
      name: 'Error401',
      component: Error401
    },
    {
      path: '/err500',
      name: 'Error500',
      component: Error500
    }

  ]
})