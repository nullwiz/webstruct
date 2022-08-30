import axios from "axios";

const url = process.env.REACT_APP_API_URL;

const arrayHelper = {
    // get array from server
    getArray: function () {
        return axios.get(url + "/api/structures/arrays?op=get");
    }
    // add Array to array
    , addArray: function (value) {
        return axios.post(url + "/api/structures/arrays?op=add&array=" + value);
    }
    , deleteArray: function () {
        return axios.post(url + "/api/structures/arrays?op=remove");
    }
    , removeDups : function () {
        return axios.post(url + "/api/structures/arrays?op=remDups");
    }
    , lowCase : function () {
        return axios.post(url + "/api/structures/arrays?op=toLowcase");
    }
}

export default arrayHelper; 
