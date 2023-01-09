import axios from 'axios'

const url = process.env.REACT_APP_API_URL

const hashMapHelper = {
  // get array from server
  getMap: function () {
    return axios.post(url + '/structures/operation', {
      Entity: 'hashmap',
      Type: 'get',
    })
  },
  // add Array to array
  addMap: function () {
    return axios.post(url + '/structures/operation', {
      Entity: 'hashmap',
      Type: 'add',
    })
  },
  setKV: function (key, value) {
    return axios.post(url + '/structures/operation', {
      Entity: 'hashmap',
      Type: 'set',
      Value: key + ':' + value,
    })
  },
  deleteMap: function () {
    return axios.post(url + '/structures/operation', {
      Entity: 'hashmap',
      Type: 'clear',
    })
  },
}

export default hashMapHelper
