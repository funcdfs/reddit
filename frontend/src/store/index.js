import { createStore } from 'vuex'

export default createStore({
  state: {
    user: {
      username: "",
      id: "",
    }
  },
  getters: {
    fullName(state) {
      return state.user.firstName + ' ' + state.user.lastName;
    }
  },
  mutations: {
  },
  actions: {
    updateUser(context) {
      
    }, 
  },
  modules: {
  }
})
