import Vue from 'vue'
import { fetchDevices, removeDevice, renameDevice, getDevice, tokenDevice} from '@/api/devices'

export default {
    namespaced: true,

    state: {
        devices: [],
        device: [],
        numberDevices: 0,
        token: ''
    },

    getters: {
        list: state => state.devices,
        get: state => state.device,
        getNumberDevices: state => state.numberDevices,
        token: state => state.token,
    },

    mutations: {
        setDevices: (state, res) => {
            Vue.set(state, 'devices', res.data)
            Vue.set(state, 'numberDevices', parseInt(res.headers['x-total-count']))
        },

        removeDevice: (state, uid) => {
            state.devices.splice(state.devices.findIndex(d => d.uid == uid), 1)
        },

        renameDevice: (state, data) => {
            state.devices = state.devices.map(i => i.uid == data.uid ? { ...i, name: data.name } : i);
        },

        setDevice: (state, data) => {
            Vue.set(state, 'device', data)
        },

        tokenDevice: (state, res) => {
            console.log(res.data.id) // eslint-disable-line no-console
            Vue.set(state, 'token', res.data.id)
        }
    },

    actions: {
        fetch: async (context, data) => {
            let res = await fetchDevices(data.per_page,data.page)
            context.commit('setDevices', res)
        },

        remove: async (context, uid) => {
            await removeDevice(uid);
            context.commit('removeDevice', uid)
        },

        rename: async (context, data) => {
            await renameDevice(data);
            context.commit('renameDevice', data)
        },

        get: async (context,uid)  => {
            let res = await getDevice(uid)
            context.commit('setDevice', res.data)
        },

        token: async (context) => {
            let res = await tokenDevice();
            context.commit('tokenDevice', res)
        },
    }
}