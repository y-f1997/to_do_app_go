<template>
<div id="error" v-html="error" v-if="error"></div>
</template>

<script>
import axios from 'axios';
export default {
  name: 'Error4xx',
  data() {
    return {
      error: null
    }
  },
  watch: {
    $route: function(){
      this.error=null;
    }
  },
  created: function(){
    var self = this;
    axios.interceptors.response.use(
      response => {
        return response;
      },
      error => {
        console.log(error)
         self.error = "<div class='alert alert-warning'>" + error.response.data.code + " :" + error.response.data.error + "</div>"
        return Promise.reject(error);
      }
    );

  }

}
</script>

<style>
#error {
  position: absolute;
  top: 16%;
  left: 50%;
  transform: translate(-50%, -50%);
  -webkit-transform: translate(-50%, -50%);
  -ms-transform: translate(-50%, -50%);
}
</style>
