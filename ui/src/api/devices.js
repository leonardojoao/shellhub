import http from '@/helpers/http'

export const
    fetchDevices = async () => {
        return http().get('/devices')
    },

    removeDevice = async (uid) => {
        return http().delete(`/devices/${uid}`)
    },

    renameDevice = async (data) => {
        return http().patch(`/devices/${data.uid}`, { name: data.name })
    },

    getDevice = async (uid) => {
        return http().get(`/devices/${uid}`)
    },

    tokenDevice = async () => {
        return http().get('/token/new')
    }