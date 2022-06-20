import { createStore } from 'vuex'
import { ModuleUser } from './user'

export default createStore({
	state: {
	},
	getters: {
	},
	mutations: {
		// mutations 中不可以实现异步操作，实现简单的数据修改

	},
	actions: {
		// 复杂的修改放在 actions 中，
	},
	modules: {
		// modules 可以对 state 进行模块的分割 
		user: ModuleUser,
	}
})
