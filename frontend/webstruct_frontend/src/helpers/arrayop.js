import axios from 'axios'

const url = process.env.REACT_APP_API_URL

const arrayHelper = {
  // get array from server
  getArray: function () {
    return axios.get(url + '/structures/operation')
  },
  // add Array to array
  addArray: function (value) {
    return axios.post(url + '/structures/operation', {
      Entity: 'string',
      Type: 'add',
      Value: value,
    })
  },
  deleteArray: function () {
    return axios.post(url + '/structures/operation')
  },
  removeDups: function () {
    return axios.post(url + '/structures/operation', {
      Entity: 'string',
      Type: 'removeDups',
    })
  },
  isPalindrome: function () {
    return axios.post(url + '/structures/operation', {
      Entity: 'string',
      Type: 'isPalindrome',
    })
  },
  lowCase: function () {
    return axios.post(url + '/structures/operation', {
      Entity: 'string',
      Type: 'lowCase',
    })
  },
}

export default arrayHelper
